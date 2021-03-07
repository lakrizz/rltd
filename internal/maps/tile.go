package maps

import (
	"image/color"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten"
	"github.com/lakrizz/rltd/internal/env"
)

const (
	TILE_EMPTY = iota
	TILE_START
	TILE_END
	TILE_PATH
	TILE_OFFENSIVE
	TILE_SUPPORT
	TILE_DEFENSIVE
)

type Tile struct {
	Id      int
	x, y    float64
	Type    int
	image   *ebiten.Image
	options *ebiten.DrawImageOptions
}

func (t *Tile) Init() error {
	c := gg.NewContext(env.TileWidth, env.TileHeight)
	c.DrawRectangle(0, 0, 32, 32)
	switch t.Type {
	case TILE_EMPTY:
		c.SetColor(color.RGBA{0xAA, 0xAA, 0xAA, 0xFF})
		break
	case TILE_START:
		c.SetColor(color.RGBA{0xCA, 0xFF, 0x70, 0xFF})
		break
	case TILE_END:
		c.SetColor(color.RGBA{0x99, 0x32, 0xCC, 0xFF})
		break
	case TILE_PATH:
		c.SetColor(color.RGBA{0xFF, 0xDE, 0xAD, 0xFF})
		break
	}
	c.Fill()
	ei, err := ebiten.NewImageFromImage(c.Image(), ebiten.FilterDefault)
	if err != nil {
		return err
	}
	t.image = ei
	op := &ebiten.DrawImageOptions{}
	t.options = op
	tx, ty := t.options.GeoM.Apply(t.x, t.y)
	t.options.GeoM.Translate(tx, ty)
	return nil
}

func (t *Tile) Update(screen *ebiten.Image) {
}

func (t *Tile) Draw(screen *ebiten.Image) {
	screen.DrawImage(t.image, t.options)
}
