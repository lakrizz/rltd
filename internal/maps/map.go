package maps

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/lakrizz/rltd/internal/env"
	"github.com/lakrizz/rltd/pkg/generators"
)

type Map struct {
	Width  int // width in number of tiles
	Height int // height in number of tiles
	Tiles  [][]*Tile
	maze   *generators.Maze
}

func GenerateMap(maze *generators.Maze) (*Map, error) {
	m := &Map{}
	m.Width = env.MapWidth
	m.Height = env.MapHeight
	m.maze = maze
	m.Tiles = make([][]*Tile, m.Height)

	for i := 0; i < m.Height; i++ {
		m.Tiles[i] = make([]*Tile, m.Width)
	}

	return m, nil
}

func (m *Map) GenerateTiles() error {
	for yi, y := range m.maze.Tiles {
		for xi, x := range y {
			tt := &Tile{x: float64(xi * env.TileWidth), y: float64(yi * env.TileHeight), Id: yi + xi, Type: x.Type}
			tt.Init()
			m.Tiles[yi][xi] = tt
		}
	}
	return nil
}

func (m *Map) Update(img *ebiten.Image) error {
	for _, y := range m.Tiles {
		for _, x := range y {
			x.Update(img)
		}
	}
	return nil
}

func (m *Map) Draw(img *ebiten.Image) error {
	for _, y := range m.Tiles {
		for _, x := range y {
			x.Draw(img)
		}
	}
	return nil
}

func (m *Map) Init() error {
	return m.GenerateTiles()
}
