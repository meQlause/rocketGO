package game

import (
	"image/color"
	"math"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/meQlause/rocketGo/config"
	"github.com/meQlause/rocketGo/object"
	"golang.org/x/image/font/basicfont"
)

type Game struct {
    Rocket *object.Object
    Fire *object.Object
    iteration, x, y, speedX, speedY, rotateSpeed, energy, lastAngle float64
}

func NewGame(rocket []byte, fire []byte, soundByte []byte) *Game {
    return &Game{
        Rocket: object.NewObject(float64(32), rocket).LoadWav(soundByte, audio.NewContext(44100)),
        Fire: object.NewObject(float64(32), fire),
    }
}

func (g *Game) Update() error {
    g.iteration++
    v := getVelocityWithJoules(g.energy)

    g.calculateY(v, float64(1.0 / 60.0))
    g.calculateX(v, float64(1.0 / 60.0))

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
    face := text.NewGoXFace(basicfont.Face7x13)
    drawText := func(str string, x, y float64) {
        var opts text.DrawOptions
        opts.GeoM.Scale(1.5, 1.5)
        opts.GeoM.Translate(x, y)                 
        opts.ColorScale.ScaleWithColor(color.White) 
        text.Draw(screen, str, face, &opts)
    }

    drawText("Speed X: "+strconv.FormatFloat(g.speedX/(1.0/60.0), 'f', 2, 64)+" m/s",0, 860)
    drawText("Distance X: "+strconv.FormatFloat(g.x, 'f', 2, 64)+" m",0, 880)
    drawText("Speed Y: "+strconv.FormatFloat(g.speedY/(1.0/60.0), 'f', 2, 64)+" m/s",0, 900)
    drawText("Distance Y: "+strconv.FormatFloat(g.y, 'f', 2, 64)+" m",0, 920)
    drawText("Energy (Joules): "+strconv.FormatFloat(g.energy, 'f', 2, 64)+" J",0, 940)
    drawText("Rocket's mass (KG): 3.000.000 KG",0, 960)
    drawText("Sin: "+strconv.FormatFloat(math.Sin(g.Rocket.Angle), 'f', 2, 64),0, 980)
    drawText("Cos: "+strconv.FormatFloat(math.Cos(g.Rocket.Angle), 'f', 2, 64),0, 1000)
    drawText("Gravity Acceleration (Earth): 9.8 m/s²",0, 1020)
    
    g.Rocket.Draw(screen, nil)
    if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
        g.Rocket.PlayFire()
        g.Fire.Draw(screen, g.Rocket)
    } else {
        g.Rocket.StopFire()
    }
}

func (g *Game) calculateY(v float64, relativeTime float64) {
    isGoingUp := math.Cos(g.Rocket.Angle) < 0              
    if isGoingUp {
        g.speedY += ((9.8 * relativeTime) - (v * -1)) * relativeTime
    }  else {
        g.speedY += ((9.8 * relativeTime) - v) * relativeTime
    }
}

func (g *Game) calculateX(v float64, relativeTime float64) {
    isGoingRight := math.Sin(g.Rocket.Angle) > 0 
    if isGoingRight {
        g.speedX += v * relativeTime
    } else {
        g.speedX -= v * relativeTime
    }
}
    
func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
    return config.ScreenWidth, config.ScreenHeight
}

func getVelocityWithJoules(joules float64) float64 {
    return math.Sqrt((2*joules)/3000000)
}