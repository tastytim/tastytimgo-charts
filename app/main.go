package main

import (
	"github.com/tastytim/charts"
	"image/color"
	"image/png"
	"os"
)

func main() {
	chart := NewBulletChart(600, 60, 70, 0, 10, color.RGBA{240, 240, 240, 255})
	chart.addLabels()

	baseBand := charts.BaseBand{10, 10, charts.Red}
	firstBand := charts.Band{18, 10, 26, 10, charts.green}
	targetBand := charts.Band{0, 25, 18, 25, charts.blue}

	chart.AddBaseBand(&baseBand)
	chart.AddBand(&firstBand)
	chart.AddBand(&targetBand)

	f, err := os.Create("bullet_chart.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, chart.Image)
}
