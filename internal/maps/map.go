package maps

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/lakrizz/rltd/internal/env"
)

type Map struct {
	Width  int // width in number of tiles
	Height int // height in number of tiles
	Tiles  [][]*Tile
}

func GenerateMap(seed int) (*Map, error) {
	m := &Map{}
	m.Width = 8
	m.Height = 8
	return m, nil
}

func (m *Map) GenerateTiles() error {
	m.Tiles = make([][]*Tile, m.Height)
	for y := 0; y < m.Height; y++ {
		m.Tiles[y] = make([]*Tile, m.Width)
		for x := 0; x < m.Width; x++ {
			tt := &Tile{x: float64(x * env.TileWidth), y: float64(y * env.TileHeight)}
			tt.Id = y + x
			tt.Init()
			m.Tiles[y][x] = tt
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
