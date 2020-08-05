package main

import (
	"fmt"
	"math"

	"github.com/Arafatk/glot"
	"github.com/jfcg/butter"
)

func main() {
	// we will mix 200 HZ and 2000 HZ signal
	f1 := 20.  //
	f2 := 500. // we will apply a 500 hz~1200hz band pass  to filter 500HZ out
	f3 := 2000.

	// use 5000 HZ as sample rate
	fs := 5000.0
	dt := 1. / fs // interval calculated due to sample rate

	fc := 100.0 // we will design a 500 hz~1200hz band pass filter
	fd := 1200.0
	wc := 2 * math.Pi * fc / fs // normalize frequency
	wd := 2 * math.Pi * fd / fs // normalize frequency

	LPF := butter.NewBandPass2(wc, wd) // initialize 2nd order band pass butterworth filter

	n := 100 // plot 100 points for demonstrating
	t := .0  // initalize time axis

	var tPoints []float64
	var uPoints []float64
	var y1Points []float64 // y1Points are those after filtered
	var y2Points []float64 // will filter a second time to get y2Points

	// signals came out after this band pass butterworth filter
	for i := 0; i < n; i++ {
		t += dt
		tPoints = append(tPoints, t)
		// mix a 200 hz and a 2000 hz
		ulow := math.Sin(2 * math.Pi * f1 * t)
		umiddle := math.Sin(2 * math.Pi * f2 * t)
		uhigh := 0.5 * math.Sin(2*math.Pi*f3*t)

		u := ulow + umiddle + uhigh
		uPoints = append(uPoints, u)

		y1 := LPF.Next(u)
		y1Points = append(y1Points, y1)
	}

	// take the output of above as new input, filter it the 2nd time

	for i := 0; i < n; i++ {

		y2 := LPF.Next(y1Points[i])
		y2Points = append(y2Points, y2)
	}

	plot2D(tPoints, uPoints, "plot of u", "plot-u.jpg")
	plot2D(tPoints, y1Points, "plot of y1 :filtered once", "plot-y1.jpg")
	plot2D(tPoints, y2Points, "plot of y2: filtered twice", "plot-y2.jpg")

}

func plot2D(xPoints, yPoints []float64, title string, figName string) {
	dimensions := 2
	persist := false
	debug := false
	plot, _ := glot.NewPlot(dimensions, persist, debug)

	if len(yPoints) != len(xPoints) {
		fmt.Println("len(yPoints) != len(xPoints)")
		return
	}

	points := [][]float64{xPoints, yPoints}

	pointGroupName := title
	style := "lines"
	plot.AddPointGroup(pointGroupName, style, points)
	plot.SetTitle(title)
	plot.SetXLabel("")
	plot.SetYLabel("")
	plot.SavePlot(figName)
}
