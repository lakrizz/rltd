package maps

import (
	"image/color"

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
			tt := &Tile{x: xi * env.TileWidth, y: yi * env.TileHeight, Id: yi + xi, Type: x.Type}
			tt.Init()
			m.Tiles[yi][xi] = tt
		}
	}
	return nil
}

func (m *Map) Update(img *ebiten.Image) error {
	mx, my := ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		xx := mx / env.TileWidth
		yy := my / env.TileHeight
		if xx >= 0 && xx < env.MapWidth && yy >= 0 && yy < env.MapHeight {

			t := m.Tiles[yy][xx]
			if t.Type == TILE_EMPTY {
				t.Changes = append(t.Changes, func(t *Tile) {
					t.gg.Clear()
					t.gg.SetColor(color.RGBA{0x00, 0xDE, 0xAD, 0xFF})
					t.gg.Fill()
					ei, _ := ebiten.NewImageFromImage(t.gg.Image(), ebiten.FilterDefault)
					t.image = ei
				})
			}
		}

		for _, y := range m.Tiles {
			for _, x := range y {
				x.Update(img)
			}
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
