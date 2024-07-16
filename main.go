package main

import (
	"log"
	"os"
	"test/symbols"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

// Game structure
type Game struct {
	interpreter *interp.Interpreter
}

// Update method
func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		reloadScripts()
	}
	return nil
}

// Draw method
func (g *Game) Draw(screen *ebiten.Image) {
	drawDebugString(screen)
}

// Layout method
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

var drawDebugString func(*ebiten.Image)

func reloadScripts() {
	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)
	i.Use(symbols.Symbols)

	bytes, err := os.ReadFile("src/debug.go")
	if err != nil {
		log.Fatal(err)
		return
	}

	src := string(bytes)

	_, err = i.Eval(src)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Retrieve the DrawDebugString function from the script
	v, err := i.Eval("src.DrawDebugString")
	if err != nil {
		log.Fatal(err)
	}

	// Assert the function type and call it
	drawDebugString = v.Interface().(func(*ebiten.Image))
}

func main() {
	// Create the game instance
	game := &Game{
		interpreter: interp.New(interp.Options{}),
	}

	reloadScripts()

	// Set window size
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ebiten with Yaegi")

	// Run the game
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
