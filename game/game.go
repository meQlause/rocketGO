package game

import (
	"image/color"
	"math"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/meQlause/rocketGo/config"
	"github.com/meQlause/rocketGo/rocket"
	"golang.org/x/image/font/basicfont"
)

type Game struct {
    Rocket *rocket.Rocket
    iteration float64
    distance float64
    speed float64
}

func NewGame() *Game {
    return &Game{
        Rocket: rocket.NewRocket(),
        iteration: 0,
        distance: 0,
        speed: 0,
    }
}

func (g *Game) Update() error {
    g.iteration++
    g.speed = float64((g.iteration / 1000 * 9.8) * 100)
    g.distance = float64((0.5 * 9.8 * math.Pow((g.iteration/1000), 2)) * 100)
    g.Rocket.Update(&g.distance)
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    text.Draw(screen, "Speed : "+strconv.FormatFloat(g.speed, 'f', 2, 64)+" m/s",
        basicfont.Face7x13, 10, 100, color.White)

    text.Draw(screen, "Distance : "+strconv.FormatFloat(g.distance, 'f', 2, 64)+" m",
        basicfont.Face7x13, 10, 120, color.White)
    g.Rocket.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return config.ScreenWidth, config.ScreenHeight
}
