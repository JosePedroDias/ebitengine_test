package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"

	raudio "ebitengine_test/audio"
	"ebitengine_test/fonts"
	"ebitengine_test/images"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	screenWidth  = 320
	screenHeight = 240

	frameOX     = 0
	frameOY     = 32
	frameWidth  = 32
	frameHeight = 32
	frameCount  = 8
)

var (
	runnerImage    *ebiten.Image
	silkscreenFont *text.GoTextFaceSource
	audioContext   *audio.Context
	jab8Player     *audio.Player
)

type Game struct {
	count int
	dx    float64
	x     float64
}

func NewGame() ebiten.Game {
	g := &Game{}
	g.init()
	return g
}

func (g *Game) Update() error {
	// TODO cumbersome logic?
	if inpututil.IsKeyJustReleased(ebiten.KeyArrowLeft) || inpututil.IsKeyJustReleased(ebiten.KeyArrowRight) {
		g.dx = 0
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		g.dx = -1
		jab8Player.Rewind()
		jab8Player.Play()
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		g.dx = 1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return bytes.ErrTooLarge
	}

	g.x += g.dx
	g.count++

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//ebitenutil.DebugPrint(screen, "Hello, World!")

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(frameWidth)/2+g.x, -float64(frameHeight)/2)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	i := (g.count / 5) % frameCount
	sx, sy := frameOX+i*frameWidth, frameOY
	screen.DrawImage(runnerImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)

	msg := fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS())
	op2 := &text.DrawOptions{}
	//op2.GeoM.Translate(0, 20)
	//op2.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, msg, &text.GoTextFace{
		Source: silkscreenFont,
		Size:   12,
	}, op2)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) init() {
	if audioContext == nil {
		audioContext = audio.NewContext(48000)

		jab8, err := wav.DecodeF32(bytes.NewReader(raudio.Jab8_wav))
		if err != nil {
			log.Fatal(err)
		}

		jab8Player, err = audioContext.NewPlayerF32(jab8)
		if err != nil {
			log.Fatal(err)
		}
	}

	font, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.Silkscreen_regular))
	if err != nil {
		log.Fatal(err)
	}
	silkscreenFont = font

	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	if err != nil {
		log.Fatal(err)
	}
	runnerImage = ebiten.NewImageFromImage(img)
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
