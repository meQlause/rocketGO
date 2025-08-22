package main

import (
	_ "embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/meQlause/rocketGo/config"
	"github.com/meQlause/rocketGo/game"
)

//go:embed assets/rocket_sound.wav
var rocketSoundBytes []byte

//go:embed assets/rocket.png
var rocket []byte

//go:embed assets/fire.png
var fire []byte

func main() {
    g := game.NewGame(rocket, fire, rocketSoundBytes)

    ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)
    ebiten.SetWindowTitle("RocketGO")

    if err := ebiten.RunGame(g); err != nil {
        log.Fatal(err)
    }
}
