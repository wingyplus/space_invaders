package main

import (
	"image/color"
	"log"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type Game struct {
	window   *sdl.Window
	renderer *sdl.Renderer

	width, height int

	running bool

	lastTimeElasped int64
	dt              int64

	invaders *InvaderContainer
}

func (g *Game) Running() bool {
	return g.running
}

// Init everything that used in this game.
func (g *Game) Init() {
	g.width, g.height = 800, 600

	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow("Space Invaders", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, g.width, g.height, sdl.WINDOW_SHOWN)
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

	g.invaders = NewInvaderContainer(
		g,
		[][]*Invader{
			{A(), A(), A(), A(), A(), A(), B(), B(), A(), A(), A()},
		},
	)

	// initialize delta time
	g.calculateDeltaTime()
}

func (g *Game) Width() int {
	return g.width
}

// Update all the component tree and re-calculate delta time in each frame
func (g *Game) Update() {
	g.invaders.Update(g.dt)
	g.calculateDeltaTime()
}

func (g *Game) calculateDeltaTime() {
	ticks := time.Now().UnixNano()
	g.dt = ticks - g.lastTimeElasped
	g.lastTimeElasped = ticks
}

// Render all of graphics in nodes
func (g *Game) Render() {
	drawBackgroundColor(g.renderer, color.Black)

	// draw an aliens
	g.invaders.Render(g.renderer)

	// display to screen
	g.renderer.Present()
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
}
