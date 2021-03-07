package main

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/lakrizz/rltd/internal/env"
	"github.com/lakrizz/rltd/internal/interfaces"
	"github.com/lakrizz/rltd/internal/maps"
	"github.com/lakrizz/rltd/pkg/generators"
)

// Game implements ebiten.Game interface.
type Game struct {
	Objects []interfaces.Renderable
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update(img *ebiten.Image) error {
	// Write your game's logical update.
	for _, v := range g.Objects {
		v.Update(img)
	}
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	for _, v := range g.Objects {
		v.Draw(screen)
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}

func main() {
	rand.Seed(time.Now().UnixNano())
	maze := generators.GenerateMaze(env.MapWidth, env.MapHeight, rand.Int63())

	game := &Game{}
	game.Objects = make([]interfaces.Renderable, 0)
	m, err := maps.GenerateMap(maze)
	if err != nil {
		panic(err)
	}

	game.Objects = append(game.Objects, m)
	// Sepcify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Your game's title")
	// Call ebiten.RunGame to start your game loop.

	for _, v := range game.Objects {
		if v.Init() != nil {
			panic(errors.New("something aua"))
		}
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
