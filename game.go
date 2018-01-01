package main

import (
	"image/color"
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

type Game struct {
	window   *sdl.Window
	renderer *sdl.Renderer

	running bool
}

func (g *Game) Running() bool {
	return g.running
}

// Init everything that used in this game.
func (g *Game) Init() {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow("Space Invaders", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		panic(err)
	}

	g.window = window
	g.renderer = renderer

	g.running = true
}

func (g *Game) Render() {
	drawBackgroundColor(g.renderer, color.Black)
}

func (g *Game) HandleEvent() {
	evt := sdl.PollEvent()
	if evt == nil {
		return
	}

	switch evt.(type) {
	case *sdl.QuitEvent:
		g.running = false
	}
}

func (g *Game) Cleanup() {
	g.renderer.Destroy()
	g.window.Destroy()
	sdl.Quit()
}

func NewGame() *Game {
	return &Game{}
}

func drawBackgroundColor(renderer *sdl.Renderer, color color.Color) {
	r, g, b, a := color.RGBA()
	if err := renderer.SetDrawColor(uint8(r), uint8(g), uint8(b), uint8(a)); err != nil {
		log.Fatal(err)
	}
	renderer.Clear()
	renderer.Present()
}
