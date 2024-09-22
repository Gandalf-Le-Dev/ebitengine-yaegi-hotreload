package src

import (
	"fmt"
	"image/color"

	"github.com/Gandalf-Le-Dev/ebiten-yaegi/engine"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/exp/rand"
)

var mainColor = color.White

func Init(g *engine.GameStruct) error {
	resetBall(g)
	return nil
}

func Update(g *engine.GameStruct) error {
	// Move paddles
	if ebiten.IsKeyPressed(ebiten.KeyW) && g.Player1Y > 0 {
		g.Player1Y -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) && g.Player1Y < engine.ScreenHeight-engine.PaddleHeight {
		g.Player1Y += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) && g.Player2Y > 0 {
		g.Player2Y -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) && g.Player2Y < engine.ScreenHeight-engine.PaddleHeight {
		g.Player2Y += 5
	}

	// Move ball
	g.BallX += g.BallDX
	g.BallY += g.BallDY

	// Ball collision with top and bottom
	if g.BallY <= 0 || g.BallY >= engine.ScreenHeight-engine.BallSize {
		g.BallDY = -g.BallDY
	}

	// Ball collision with paddles
	if g.BallX <= engine.PaddleWidth && g.BallY+engine.BallSize >= g.Player1Y && g.BallY <= g.Player1Y+engine.PaddleHeight {
		g.BallDX = -g.BallDX
	}
	if g.BallX >= engine.ScreenWidth-engine.PaddleWidth-engine.BallSize && g.BallY+engine.BallSize >= g.Player2Y && g.BallY <= g.Player2Y+engine.PaddleHeight {
		g.BallDX = -g.BallDX
	}

	// Score
	if g.BallX <= 0 {
		g.Score2++
		resetBall(g)
	}
	if g.BallX >= engine.ScreenWidth-engine.BallSize {
		g.Score1++
		resetBall(g)
	}

	return nil
}

func Draw(g *engine.GameStruct, screen *ebiten.Image) error {
	// Draw paddles
	vector.DrawFilledRect(screen, 0, g.Player1Y, engine.PaddleWidth, engine.PaddleHeight, mainColor, false)
	vector.DrawFilledRect(screen, engine.ScreenWidth-engine.PaddleWidth, g.Player2Y, engine.PaddleWidth, engine.PaddleHeight, mainColor, false)

	// Draw ball
	vector.DrawFilledCircle(screen, g.BallX+engine.BallSize/2, g.BallY+engine.BallSize/2, engine.BallSize/2, mainColor, false)

	// Draw score
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Player 1: %d | Player 2: %d", g.Score1, g.Score2), 0, 20)

	return nil
}

func resetBall(g *engine.GameStruct) {
	g.BallX = engine.ScreenWidth / 2
	g.BallY = engine.ScreenHeight / 2
	g.BallDX = 4 * float32(rand.Intn(2)*2-1)
	g.BallDY = 4 * float32(rand.Intn(2)*2-1)
}
