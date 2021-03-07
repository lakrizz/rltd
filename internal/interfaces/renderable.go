package interfaces

import "github.com/hajimehoshi/ebiten"

type Renderable interface {
	Init() error
	Update(*ebiten.Image) error
	Draw(*ebiten.Image) error
}
