package engine

import (
	"embed"
	"log"

	"github.com/Gandalf-Le-Dev/ebiten-yaegi/engineUtils"
	"github.com/Gandalf-Le-Dev/ebiten-yaegi/interpreterUtils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/traefik/yaegi/interp"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
	PaddleWidth  = 10
	PaddleHeight = 60
	BallSize     = 10
)

type GameStruct struct {
	Src      *embed.FS
	Player1Y float32
	Player2Y float32
	BallX    float32
	BallY    float32
	BallDX   float32
	BallDY   float32
	Score1   int
	Score2   int
}

var initMethod func(game *GameStruct) error
var updateMethod func(game *GameStruct) error
var drawMethod func(game *GameStruct, screen *ebiten.Image) error
var game *GameStruct
var interpreter *interp.Interpreter

func InitEngine(src *embed.FS) {
	interpreter = interpreterUtils.NewInterpreter()
	game = &GameStruct{
		Src:      src,
		Player1Y: ScreenHeight/2 - PaddleHeight/2,
		Player2Y: ScreenHeight/2 - PaddleHeight/2,
	}

	game.LoadScripts()

	// Set window size
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Ebiten with Yaegi")

	// Run the game
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

	initMethod(game)
}

// Update method
func (g *GameStruct) Update() error {
	if inpututil.IsKeyJustReleased(ebiten.KeyR) {
		g.LoadScripts()
	}

	updateMethod(g)

	return nil
}

// Draw method
func (g *GameStruct) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Press R to reload scripts")
	drawMethod(g, screen)
}

// Layout method
func (g *GameStruct) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func GetScreenSize() (int, int) {
	return ScreenWidth, ScreenHeight
}

func (g *GameStruct) LoadScripts() {
	// Load all scripts
	files, err := engineUtils.FindFilesRecursive(g.Src)
	if err != nil {
		log.Default().Printf("Error loading scripts: %v\n", err)
	}

	for _, file := range files {
		log.Default().Printf("Loading file: %s\n", file)
		content, err := engineUtils.LoadFile(file)
		if err != nil {
			log.Default().Printf("Error loading file: %v\n", err)
		}

		_, err = interpreter.Eval(content)
		if err != nil {
			log.Default().Printf("Error evaluating file: %v\n", err)
		}
	}

	v, err := interpreter.Eval("src.Init")
	if err != nil {
		log.Fatal(err)
	}

	// Assert the function type and call it
	initMethod = v.Interface().(func(*GameStruct) error)

	v, err = interpreter.Eval("src.Update")
	if err != nil {
		log.Fatal(err)
	}

	// Assert the function type and call it
	updateMethod = v.Interface().(func(*GameStruct) error)

	v, err = interpreter.Eval("src.Draw")
	if err != nil {
		log.Fatal(err)
	}

	drawMethod = v.Interface().(func(*GameStruct, *ebiten.Image) error)
}
