package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/meQlause/rocketGo/config"
	"github.com/meQlause/rocketGo/game"
)

func main() {
    g := game.NewGame()

    ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)
    ebiten.SetWindowTitle("Simple Ebiten Game")

    if err := ebiten.RunGame(g); err != nil {
        log.Fatal(err)
    }
}
