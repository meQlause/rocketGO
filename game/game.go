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
    rotateSpeed float64
    energy float64
}

func NewGame() *Game {
    return &Game{
        Rocket: rocket.NewRocket(),
        iteration: 0,
        distance: 0,
        speed: 0,
        rotateSpeed: 0,
        energy: 0,
    }
}

func (g *Game) Update() error {
    g.iteration++

    speedCanceled := math.Sqrt((2*g.energy)/30000000)
    dt := 1.0 / 60.0              
    g.speed += (9.8 * dt) - speedCanceled             
    g.distance += g.speed * dt     
    g.Rocket.Update(&g.distance)
    
    if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
        g.rotateSpeed += 0.001
    }
        
    if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
        g.rotateSpeed -= 0.001
    }

    if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
        g.iteration += 2
        g.energy += 10000
    } else {
        g.energy = 0
    }

    g.Rocket.Angle += g.rotateSpeed
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    text.Draw(screen, "Speed : "+strconv.FormatFloat(g.speed, 'f', 2, 64)+" m/s",
        basicfont.Face7x13, 10, 100, color.White)
    text.Draw(screen, "Distance : "+strconv.FormatFloat(g.distance, 'f', 2, 64)+" m",
        basicfont.Face7x13, 10, 120, color.White)
    text.Draw(screen, "Energy (Joule) : "+strconv.FormatFloat(g.energy, 'f', 2, 64)+" J",
        basicfont.Face7x13, 10, 140, color.White)
    text.Draw(screen, "Rocket's mass (KG) : 30.000.000 KG",
        basicfont.Face7x13, 10, 160, color.White)
    g.Rocket.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return config.ScreenWidth, config.ScreenHeight
}
