package main

import (
	"embed"

	"github.com/Gandalf-Le-Dev/ebiten-yaegi/engine"
)

//go:embed src/*
var src embed.FS

func main() {
	engine.InitEngine(&src)
}
