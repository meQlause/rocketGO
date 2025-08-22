package game

import (
	"image/color"
	"math"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/meQlause/rocketGo/config"
	"github.com/meQlause/rocketGo/object"
	"golang.org/x/image/font/basicfont"
)

type Game struct {
    Rocket *object.Object
    Fire *object.Object
    iteration, x, y, speedX, speedY, rotateSpeed, energy, lastAngle float64
}

func NewGame() *Game {
    return &Game{
        Rocket: object.NewObject(float64(32), "assets/rocket.png").LoadWav("assets/rocket_sound.wav", audio.NewContext(44100)),
        Fire: object.NewObject(float64(32), "assets/fire.png"),
    }
}

func (g *Game) Update() error {
    g.iteration++
    thurst := getVelocityWithJoules(g.energy)

    g.calculateY(thurst, float64(1.0 / 60.0))
    g.calculateX(thurst, float64(1.0 / 60.0))

    g.Fire.Angle = -0.55 
    if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
        g.Fire.Angle = -0.65 
        g.rotateSpeed += 0.0001
    }
        
    if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
        g.Fire.Angle = -0.45 
        g.rotateSpeed -= 0.0001
    }

    if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
        g.iteration += 2
        g.energy += 900
        g.lastAngle = math.Cos(g.Rocket.Angle - math.Pi/2)
    } else {
        g.energy = 0
    }
    
    g.x += g.speedX
    g.y += g.speedY     
    g.Rocket.Update(g.x, g.y)
    g.Fire.Update(g.x, g.y)
    g.Rocket.Angle += g.rotateSpeed
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    text.Draw(screen, "Speed X: "+strconv.FormatFloat(g.speedX / float64(1.0 / 60.0) , 'f', 2, 64)+" m/s",
        basicfont.Face7x13, 10, 860, color.White)
    text.Draw(screen, "Distance X: "+strconv.FormatFloat(g.x, 'f', 2, 64)+" m",
        basicfont.Face7x13, 10, 880, color.White)
    text.Draw(screen, "Speed Y: "+strconv.FormatFloat(g.speedY / float64(1.0 / 60.0), 'f', 2, 64)+" m/s",
        basicfont.Face7x13, 10, 900, color.White)
    text.Draw(screen, "Distance Y: "+strconv.FormatFloat(g.y, 'f', 2, 64)+" m",
        basicfont.Face7x13, 10, 920, color.White)
    text.Draw(screen, "Energy (Joules) : "+strconv.FormatFloat(g.energy, 'f', 2, 64)+" J",
        basicfont.Face7x13, 10, 940, color.White)
    text.Draw(screen, "Rocket's mass (KG) : 3.000.000 KG",
        basicfont.Face7x13, 10, 960, color.White)
    text.Draw(screen, "Sin : "+strconv.FormatFloat(math.Sin(g.Rocket.Angle), 'f', 2, 64),
        basicfont.Face7x13, 10, 980, color.White)
    text.Draw(screen, "Cos : "+strconv.FormatFloat(math.Cos(g.Rocket.Angle), 'f', 2, 64),
        basicfont.Face7x13, 10, 1000, color.White)
    text.Draw(screen, "Gratvity Acceleration (Earth) : 9.8 m/s",
        basicfont.Face7x13, 10, 1020, color.White)
    g.Rocket.Draw(screen, nil)
    if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
        g.Rocket.PlayFire()
        g.Fire.Draw(screen, g.Rocket)
    } else {
        g.Rocket.StopFire()
    }
}

func (g *Game) calculateY(thurst float64, relativeTime float64) {
        isGoingUp := math.Cos(g.Rocket.Angle) < 0              
    if isGoingUp {
        g.speedY += ((9.8 * relativeTime) - (thurst * -1)) * relativeTime
    }  else {
        g.speedY += ((9.8 * relativeTime) - thurst) * relativeTime
    }
}

func (g *Game) calculateX(thurst float64, relativeTime float64) {
    isGoingRight := math.Sin(g.Rocket.Angle) > 0 
    if isGoingRight {
        g.speedX += thurst * relativeTime
    } else {
        g.speedX -= thurst * relativeTime
    }
}
    
func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
    return config.ScreenWidth, config.ScreenHeight
}

func getVelocityWithJoules(joules float64) float64 {
    return math.Sqrt((2*joules)/3000000)
}