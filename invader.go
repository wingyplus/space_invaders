package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	space   = 10
	padding = 10
)

const (
	invaderWidth  = 32
	invaderHeight = 32
)

type InvaderContainer struct {
	Grid [][]*Invader
	Game *Game

	// begin position of grid
	x, y             int32
	xoffset, yoffset int32
	direction        MoveDirection

	interval time.Duration
}

func (container *InvaderContainer) updateInterval(dt int64) {
	container.interval += time.Duration(dt)
}

func (container *InvaderContainer) resetInterval() {
	container.interval = 0
}

// Update detect grid position and decide should be move next.
func (container *InvaderContainer) Update(dt int64) {
	// move invaders every ~0.5 seconds
	if container.interval <= 500*time.Millisecond {
		container.updateInterval(dt)
		return
	}

	container.resetInterval()

	for i := 0; i < len(container.Grid); i++ { // row
		for j := 0; j < len(container.Grid[i]); j++ { // column
			invader := container.Grid[i][j]
			// TODO: cleanup calculation movement.
			x, y := container.x+padding+((invader.width+space)*int32(j)), container.y+10
			invader.SetPos(x, y)

			if j == len(container.Grid[i])-1 {
				container.xoffset = invader.X() + invaderWidth + 10
			}
		}
	}

	switch container.direction {
	case Right:
		if int(container.xoffset) < container.Game.Width() {
			// move to the right side
			container.x += 10
		} else {
			container.direction = DownThenLeft
		}
	case Left:
		if int(container.x) > 0 {
			container.x -= 10
		} else {
			container.direction = DownThenRight
		}
	case DownThenLeft, DownThenRight:
		container.y += 10 + invaderHeight
		if container.direction == DownThenLeft {
			container.direction = Left
		} else if container.direction == DownThenRight {
			container.direction = Right
		}
	}
}

func (container *InvaderContainer) Render(renderer *sdl.Renderer) {
	for i := 0; i < len(container.Grid); i++ {
		for j := 0; j < len(container.Grid[i]); j++ {
			invader := container.Grid[i][j]
			invader.Render(renderer)
		}
	}
}

func NewInvaderContainer(game *Game, invaders [][]*Invader) *InvaderContainer {
	return &InvaderContainer{
		Game:      game,
		Grid:      invaders,
		direction: Right,
		x:         10,
		y:         10,
	}
}

type MoveDirection int

const (
	Left MoveDirection = iota
	Right
	DownThenLeft
	DownThenRight
)

type InvaderType int

const (
	TypeA InvaderType = iota
	TypeB
	TypeC
)

type Invader struct {
	x, y          int32
	t             InvaderType
	width, height int32
}

func (invader *Invader) SetPos(x, y int32) {
	invader.x, invader.y = x, y
}

func (invader *Invader) X() int32 {
	return invader.x
}

func (invader *Invader) Y() int32 {
	return invader.y
}

func (invader *Invader) Render(renderer *sdl.Renderer) {
	switch invader.t {
	case TypeA:
		renderer.SetDrawColor(255, 0, 0, 1)
	case TypeB:
		renderer.SetDrawColor(0, 255, 0, 1)
	}

	renderer.DrawRect(&sdl.Rect{
		X: invader.x,
		Y: invader.y,
		W: invader.width,
		H: invader.height,
	})
}

func NewInvader(t InvaderType) *Invader {
	return &Invader{
		t:      t,
		width:  32,
		height: 32,
	}
}

func A() *Invader {
	return NewInvader(TypeA)
}

func B() *Invader {
	return NewInvader(TypeB)
}
