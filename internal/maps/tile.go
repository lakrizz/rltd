package maps

import (
	"image/color"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten"
	"github.com/lakrizz/rltd/internal/env"
)

type Tile struct {
	Id      int
	x, y    float64
	image   *ebiten.Image
	options *ebiten.DrawImageOptions
}

func (t *Tile) Init() error {
	c := gg.NewContext(env.TileWidth, env.TileHeight)
	c.DrawCircle(16, 16, 16)
	if t.Id%2 == 0 {
		c.SetColor(color.White)
	} else {
		c.SetColor(color.RGBA{0xFF, 0x00, 0x00, 0xFF})
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
