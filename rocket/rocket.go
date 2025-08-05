package rocket

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Rocket struct {
    X, Y float64
    Speed float64
    Img  *ebiten.Image

}

func NewRocket() *Rocket {
    img, _, err := ebitenutil.NewImageFromFile("assets/rocket.png")
    if err != nil {
        log.Fatal(err)
    }
    return &Rocket{
        X: 400,
        Y: 0,
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
    op.GeoM.Translate(p.X, p.Y)
    screen.DrawImage(p.Img, op)
}
