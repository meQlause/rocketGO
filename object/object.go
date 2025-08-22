package object

import (
	"log"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Object struct {
    x, y, xInit, yInit, Angle, Size, Speed  float64
    Img *ebiten.Image
    audioContext *audio.Context
    objectSound *audio.Player
}


func NewObject(size float64, imageLocation string) *Object {
    if rocket, _, err := ebitenutil.NewImageFromFile(imageLocation); err == nil {
        return &Object{
            xInit: 400,
            Size: size,
            Img: rocket,
        }
    } else {
        log.Fatal(err)
    }
    return nil
}

func (p *Object) LoadWav(filename string, audioContext *audio.Context) *Object {
    p.audioContext = audioContext
    f, err := os.Open(filename)
    if err != nil {
        log.Fatalf("failed to open wav file: %v", err)
    }

    d, err := wav.DecodeWithSampleRate(44100, f)
    if err != nil {
        log.Fatalf("failed to decode wav file: %v", err)
    }

    loop := audio.NewInfiniteLoop(d, d.Length())

    if player, err := p.audioContext.NewPlayer(loop); err == nil {
        p.objectSound = player
    } else {
        log.Fatalf("failed to create player: %v", err)
    }

    return p
}


func (p *Object) Update(x float64, y float64) {
    p.x = x + p.xInit
    p.y = y + p.yInit
}

func(p *Object) PlayFire() {
    if p.objectSound != nil && !p.objectSound.IsPlaying() {
        p.objectSound.Play()
    } 
}

func (p *Object) StopFire() {
    if p.objectSound != nil {
        p.objectSound.Pause()
        if err := p.objectSound.Rewind(); err != nil {
            log.Fatal(err)
        }
    }
}

func (p *Object) Draw(screen *ebiten.Image, otherObj *Object) {
    op := &ebiten.DrawImageOptions{}
    scaleX := p.Size / float64(p.Img.Bounds().Dx())
    scaleY := p.Size / float64(p.Img.Bounds().Dy())
    op.GeoM.Scale(scaleX, scaleY)

    if otherObj == nil {
        w := float64(p.Img.Bounds().Dx()) * scaleX
        h := float64(p.Img.Bounds().Dy()) * scaleY

        op.GeoM.Translate(-w/2, -h/2)
        op.GeoM.Rotate(p.Angle)
        op.GeoM.Translate(p.x, p.y)
    } else {
        fw := float64(p.Img.Bounds().Dx()) * scaleX
        fh := float64(p.Img.Bounds().Dy()) * scaleY

        op.GeoM.Translate(-fw/2, -fh/2)
        op.GeoM.Rotate(otherObj.Angle + p.Angle)

        rh := float64(otherObj.Img.Bounds().Dy()) * (p.Size / float64(otherObj.Img.Bounds().Dy()))
        offset := rh/2 + fh/2
        dx := offset * math.Cos(otherObj.Angle + math.Pi/2)
        dy := offset * math.Sin(otherObj.Angle + math.Pi/2)

        op.GeoM.Translate(otherObj.x+dx, otherObj.y+dy)
    }

    screen.DrawImage(p.Img, op)
}

