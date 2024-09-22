package interpreterUtils

import (
	"github.com/Gandalf-Le-Dev/ebiten-yaegi/symbols"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

func NewInterpreter() *interp.Interpreter {
	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)
	i.Use(symbols.Symbols)

	return i
}
