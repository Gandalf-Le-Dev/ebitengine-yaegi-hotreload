# Ebitengine engine with hotreload thanks to yaegi

## How to use

Run the project 

`go run .`

You will see the default debug text on screen

Now you can update the debug text in the file `src/debug.go`.

Example of update:

```go
package src

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func DrawDebugString(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "THIS IS AN UPDATE")
}
```

After saving press R in game and you will see you updated text. 

 
