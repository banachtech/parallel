package main

import (
	"image/color"
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	width, height = 640, 360
	numBoids      = 500
	viewRad       = 13
	adjRate       = 0.015
)

var (
	green   = color.RGBA{10, 255, 50, 255}
	boids   [numBoids]*Boid
	boidMap [width + 1][height + 1]int
	rwlock    = sync.RWMutex{}
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, boid := range boids {
		screen.Set(int(boid.position.x+1), int(boid.position.y), green)
		screen.Set(int(boid.position.x-1), int(boid.position.y), green)
		screen.Set(int(boid.position.x), int(boid.position.y-1), green)
		screen.Set(int(boid.position.x), int(boid.position.y+1), green)
	}
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return width, height
}

func main() {
	for i, row := range boidMap {
		for j := range row {
			boidMap[i][j] = -1
		}
	}
	for i := 0; i < numBoids; i++ {
		createBoid(i)
	}
	ebiten.SetWindowSize(width*2, height*2)
	ebiten.SetWindowTitle("Boids in a Box")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
