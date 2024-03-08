package charts

import (
	"fmt"
	"image"
	"image/color"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

var (
	red   = color.RGBA{255, 16, 12, 255}
	green = color.RGBA{1, 251, 176, 255}
	blue  = color.RGBA{80, 116, 255, 255}
)

type BulletChart struct {
	Image           *image.RGBA
	Width           int        // width image
	Height          int        // height image
	XMax            int        // x axe maximum number
	XMin            int        // x axe maximum number
	Step            int        // steps on x axe
	BackGroundColor color.RGBA //color image
}

// Create new Chart
func NewBulletChart(width int, height int, xMax int, xMin int, step int, bColor color.RGBA) *BulletChart {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	return &BulletChart{Image: img, XMax: xMax, XMin: xMin, Step: step, BackGroundColor: bColor}
}

type Band struct {
	FromXStart  int
	FromXEnd    int
	FromYTop    int
	FromYBottom int
	Color       color.RGBA
}
type BaseBand struct {
	FromYTop    int
	FromYBottom int
	Color       color.RGBA
}

func (b *BulletChart) AddBand(band *Band) {
	fillRect(b.Image, image.Rect(band.FromXStart*b.Width/b.XMax, band.FromYBottom, band.FromXEnd*b.Width/b.XMax, b.Height-band.FromYBottom), band.Color)
}

func (b *BulletChart) AddBaseBand(band *BaseBand) {
	fillRect(b.Image, image.Rect(0, band.FromYBottom, b.Width, b.Height-band.FromYBottom), band.Color)
}

func (b *BulletChart) AddLabels() {
	col := color.Black
	for i := 0; i <= b.XMax; i += b.Step {
		x := i * b.Width / b.XMax
		label := fmt.Sprintf("%d", i)
		var point fixed.Point26_6
		if x == b.Width {
			point = fixed.P(x-spaceForLastLabel(b.Width), b.Height)
		} else {
			point = fixed.P(x, b.Height)
		}

		d := &font.Drawer{
			Dst:  b.Image,
			Src:  image.NewUniform(col),
			Face: basicfont.Face7x13,
			Dot:  point,
		}
		d.DrawString(label)
	}
}

func fillRect(img *image.RGBA, rect image.Rectangle, color color.Color) {
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			img.Set(x, y, color)
		}
	}
}

func spaceForLastLabel(n int) int {
	count := 0
	for n > 0 {
		n = n / 10
		count++
	}
	return count * 7
}
