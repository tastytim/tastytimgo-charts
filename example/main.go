package main

import (
	"image/color"
	"image/png"
	"os"

	"github.com/tastytim/charts/charts"
)

func main() {
	//Example bullet chart draw

	
	chart := charts.NewBulletChart(600, 60, 70, 0, 10, color.RGBA{240, 240, 240, 255})
	baseBand := charts.BaseBand{YTop: 10, YBottom: 10, Color: charts.Red}
	firstBand := charts.Band{XStart: 0, XEnd: 10, YTop: 10, YBottom: 10, Color: charts.Red}
	rangeBand := charts.Band{XStart: 10, XEnd: 35, YTop: 10, YBottom: 10, Color: charts.Green}
	targetBand := charts.Band{XStart: 0, XEnd: 18, YTop: 25, YBottom: 25, Color: charts.Blue}


	//make sure you do it in order
	chart.AddBaseBand(baseBand)
	chart.AddBand(firstBand)
	chart.AddBand(rangeBand)
	chart.AddBand(targetBand)
	chart.AddLabels()

	f, err := os.Create("bullet_chart.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, chart.Image)
}
