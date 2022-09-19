package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 680
	screenHeight = 720
	maxIter      = 512
	startCenterX = -0.286
	startCenterY = 0
	startSize    = 3
)

var (
	palette [maxIter]byte
)

func init() {
	for i := range palette {
		palette[i] = byte(math.Sqrt(float64(i)/float64(len(palette))) * 0x80)
	}
}

func color(it int) (r, g, b byte) {
	if it == maxIter {
		return 0xff, 0xff, 0xff
	}
	c := palette[it]
	return c, c, c
}

type Field struct {
	image    *ebiten.Image
	imagePix []byte
	centerX  float64
	centerY  float64
	size     float64
}

func MakeField() *Field {
	f := &Field{
		image:    ebiten.NewImage(screenWidth, screenHeight),
		imagePix: make([]byte, (screenHeight * screenWidth * 4)),
		centerX:  startCenterX,
		centerY:  startCenterY,
		size:     startSize,
	}
	f.updateField()
	return f
}

func (f *Field) updateField() {
	startTime := time.Now()
	for j := 0; j < screenHeight; j++ {
		for i := 0; i < screenWidth; i++ {
			x := float64(i)*f.size/screenWidth - f.size/2 + f.centerX
			y := (screenHeight-float64(j))*f.size/screenHeight - f.size/2 + f.centerY
			c := complex(x, y)
			z := complex(0, 0)
			it := 0
			for ; it < maxIter; it++ {
				z = z*z + c
				if real(z)*real(z)+imag(z)*imag(z) > 10 {
					break
				}
			}
			r, g, b := color(it)
			p := 4 * (i + j*screenWidth)
			f.imagePix[p] = r
			f.imagePix[p+1] = g
			f.imagePix[p+2] = b
			f.imagePix[p+3] = 0xff
		}
	}
	f.image.WritePixels(f.imagePix)
	drawTimeMill := time.Now().UnixMilli() - startTime.UnixMilli()
	fmt.Printf("f.centerX: %v\n", f.centerX)
	fmt.Printf("f.centerY: %v\n", f.centerY)
	fmt.Printf("f.size: %v\n", f.size)
	fmt.Printf("drawTimeMill: %v\n", drawTimeMill)
}

func (f *Field) Update() error {
	needRedraw := false
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		f.centerX -= 0.05 * f.size
		needRedraw = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		f.centerX += 0.05 * f.size
		needRedraw = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		f.centerY += 0.05 * f.size
		needRedraw = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		f.centerY -= 0.05 * f.size
		needRedraw = true
	}
	_, dy := ebiten.Wheel()
	if dy != 0 {
		if dy < 0 {
			f.size *= math.Pow(0.95, math.Abs(dy))
		} else {
			f.size *= math.Pow(1.05, math.Abs(dy))
		}
		needRedraw = true
	}

	if needRedraw {
		f.updateField()
	}

	return nil
}

func (g *Field) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, Mandelbrot!")
	screen.DrawImage(g.image, nil)
}

func (g *Field) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Mandelbrot")
	if err := ebiten.RunGame(MakeField()); err != nil {
		log.Fatal(err)
	}
}
