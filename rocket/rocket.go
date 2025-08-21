package rocket

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Rocket struct {
    X, Y, Angle float64
    Speed float64
    Img  *ebiten.Image
}

func NewRocket() *Rocket {
    img, _, _ := ebitenutil.NewImageFromFile("assets/rocket.png")
    return &Rocket{
        X: 400,
        Y: 0,
        Angle: 0,
        Img: img,
    }
}

func (p *Rocket) Update(distance *float64) {
    p.Y = *distance
}

func (p *Rocket) Draw(screen *ebiten.Image) {
    op := &ebiten.DrawImageOptions{}

    scaleX := 32.0 / float64(p.Img.Bounds().Dx())
    scaleY := 32.0 / float64(p.Img.Bounds().Dy())
    op.GeoM.Scale(scaleX, scaleY)

    w := float64(p.Img.Bounds().Dx()) * scaleX
    h := float64(p.Img.Bounds().Dy()) * scaleY

    op.GeoM.Translate(-w/2, -h/2)

    op.GeoM.Rotate(p.Angle)

    op.GeoM.Translate(p.X, p.Y)

    screen.DrawImage(p.Img, op)
}