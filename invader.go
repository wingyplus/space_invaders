package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type InvaderContainer struct {
	Grid [][]*Invader
	Game *Game
}

// Update detect grid position and decide should be move next.
func (container *InvaderContainer) Update(dt uint32) {
	for i := 0; i < len(container.Grid); i++ {
		for j := 0; j < len(container.Grid[i]); j++ {
			invader := container.Grid[i][j]

			if invader.direction == DownThenLeft {
				invader.direction = Left
			} else if invader.direction == DownThenRight {
				invader.direction = Right
			}

			if int(invader.X()+10) < container.Game.Width() {
				invader.direction = Right
			} else {
				invader.direction = DownThenLeft
			}
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

func A(x, y int32) *Invader {
	return &Invader{
		x:         x,
		y:         y,
		t:         TypeA,
		width:     32,
		height:    32,
		direction: Right,
	}
}
