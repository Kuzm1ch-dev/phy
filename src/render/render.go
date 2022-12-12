package render

import (
	"fmt"
	"image/color"
	"log"
	"phy/src/physic"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 640
	screenHeight = 480
	scale        = 20
)

type Game struct {
	collisionSystem *physic.CollisionSystem
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	/*
		vector.StrokeLine(screen, 100, 100, 300, 100, 1, color.RGBA{0xff, 0xff, 0xff, 0xff})
		vector.StrokeLine(screen, 50, 150, 50, 350, 1, color.RGBA{0xff, 0xff, 0x00, 0xff})
		vector.StrokeLine(screen, 50, 100+cf, 200+cf, 250, 4, color.RGBA{0x00, 0xff, 0xff, 0xff})

		vector.DrawFilledRect(screen, 50+cf, 50+cf, 100+cf, 100+cf, color.RGBA{0x80, 0x80, 0x80, 0xc0})
		vector.StrokeRect(screen, 300-cf, 50, 120, 120, 10+cf/4, color.RGBA{0x00, 0x80, 0x00, 0xff})

		vector.DrawFilledCircle(screen, 400, 400, 100, color.RGBA{0x80, 0x00, 0x80, 0x80})*/

	for _, e := range g.collisionSystem.Entities {
		switch e.Body.M_fixtureList.M_shape.GetType() {

		case 0:
			fmt.Println(fmt.Sprintf("X: %f, Y: %f, R: %f"), e.Body.GetPosition().X*scale, e.Body.GetPosition().X*scale, e.Body.M_fixtureList.M_shape.GetRadius()*scale)
			vector.StrokeCircle(screen, float32(e.Body.GetPosition().X*scale), float32(e.Body.GetPosition().X*scale), float32(e.Body.M_fixtureList.M_shape.GetRadius()*scale), 1, color.RGBA{0xff, 0x80, 0xff, 0xff})

		}
	}

	vector.StrokeCircle(screen, 400, 400, 10, 1, color.RGBA{0xff, 0x80, 0xff, 0xff})

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func Run(world *physic.CollisionSystem) {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Shapes (Ebitengine Demo)")
	if err := ebiten.RunGame(&Game{world}); err != nil {
		log.Fatal(err)
	}
}
