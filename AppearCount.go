package main

import (
	"fmt"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
)

type AppearCount map[int]int

func (input AppearCount) max() int {
	var max int
	for _, count := range input {
		if count > max {
			max = count
		}
	}
	return max
}

func (app AppearCount) print(human bool) {
	if human {
		app.iterate(10, func(counter, key, val int) {
			fmt.Printf("%2d. Count: %4d,  Appeared: %4d\n", counter, key, val)
		})
	} else {
		app.iterate(20, func(counter, key, val int) {
			fmt.Printf("%d,%d\n", key, val)
		})
	}
}

func (app AppearCount) iterate(stop int, fn func(counter, key, val int)) {
	counter := 0
	var countCond = func(counter int) bool {
		return counter < stop
	}

	max := app.max()
	for index := max; index > 0 && countCond(counter); index -= 1 {
		for key, val := range app {
			if val == index && countCond(counter) {
				counter += 1
				fn(counter, key, val)
			}
		}
	}
}

func (app AppearCount) plot(log bool, pfix string) {
	if len(app) > 0 {
		plotObj, err := plot.New()
		if err != nil {
			panic(err)
		}

		plotObj.Title.Text = "Count Appearances"
		plotObj.X.Label.Text = "appearance"
		plotObj.Y.Label.Text = "count"

		points := make(plotter.XYs, len(app))
		app.iterate(len(app), func(i, k, v int) {
			points[i-1].X = float64(v)
			points[i-1].Y = float64(k)
		})

		err = plotutil.AddScatters(plotObj, "", points)
		if err != nil {
			panic(err)
		}

		fname := "out/ac-" + pfix + ".png"
		if log {
			plotObj.X.Scale = plot.LogScale{}
			plotObj.Y.Scale = plot.LogScale{}
			fname = "out/ac-" + pfix + ".png"
		}

		// Save the plot to a PNG file.
		if err := plotObj.Save(5*vg.Inch, 5*vg.Inch, fname); err != nil {
			panic(err)
		}
	}
}
