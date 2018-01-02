package main

import (
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
	x, y      int32
	direction MoveDirection
}

// Update detect grid position and decide should be move next.
// TODO: find how to move it slowly.
func (container *InvaderContainer) Update(dt uint32) {
	for i := 0; i < len(container.Grid); i++ { // row
		for j := 0; j < len(container.Grid[i]); j++ { // column
			invader := container.Grid[i][j]
			x, y := container.x+padding+((invader.width+space)*int32(j)), container.y+10
			invader.SetPos(x, y)
		}
	}

	switch container.direction {
	case Right:
		// TODO: refactor this calculation, it's look hard to understand
		if int(container.x+padding+((invaderWidth+space)*11)+10) < container.Game.Width() {
			// move to the right side
			container.x += 10
		} else {
			container.direction = DownThenLeft
		}
	case Left:
		if int(container.x) > container.Game.Width() {
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

	direction MoveDirection
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
	renderer.SetDrawColor(255, 0, 0, 1)
	rect := sdl.Rect{
		X: invader.x,
		Y: invader.y,
		W: invader.width,
		H: invader.height,
	}

	switch invader.direction {
	case Right:
		invader.x += 10
		rect.X = invader.x
	case DownThenLeft:
		invader.y += 10
		rect.Y = invader.y
	}

	renderer.DrawRect(&rect)
}

func A() *Invader {
	return &Invader{
		t:         TypeA,
		width:     32,
		height:    32,
		direction: Right,
	}
}
