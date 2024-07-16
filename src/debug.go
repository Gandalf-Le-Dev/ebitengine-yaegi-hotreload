package src

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func DrawDebugString(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "THIS IS AN UPDATE")
}
