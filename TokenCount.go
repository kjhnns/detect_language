package main

import (
	"fmt"
	"strings"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
)

type TokenCount map[string]int

func wordCount(input []string) TokenCount {
	out := map[string]int{}
	for _, tok := range input {
		out[tok] += 1
	}
	return out
}

func runeCount(input []string) TokenCount {
	out := map[string]int{}
	for _, tok := range input {
		strings.Map(func(r rune) rune {
			out[string(r)] += 1
			return r
		}, tok)
	}
	return out
}

func pairCount(input []string) TokenCount {
	out := map[string]int{}
	var prev rune
	for _, tok := range input {
		strings.Map(func(curr rune) rune {
			if prev > 0 {
				out[string(prev)+string(curr)] += 1
			}
			prev = curr
			return curr
		}, tok)
		prev = 0
	}
	return out
}

func (input TokenCount) max() int {
	var max int
	for _, count := range input {
		if count > max {
			max = count
		}
	}
	return max
}

func (input TokenCount) appearances() AppearCount {
	out := map[int]int{}
	for _, val := range input {
		out[val] += 1
	}
	return out
}

func (tokc TokenCount) iterate(stop int, fn func(counter int, key string, val int)) {
	counter := 0
	var countCond = func(counter int) bool {
		return counter < stop
	}

	max := tokc.max()
	for index := max; index > 0 && countCond(counter); index -= 1 {
		for key, val := range tokc {
			if val == index && countCond(counter) {
				counter += 1
				fn(counter, key, val)
			}
		}
	}
}

func (out TokenCount) print(human bool) {
	total := out.sum()

	if human {
		out.iterate(10, func(counter int, key string, val int) {
			rel := float64(val) / float64(total) * 100
			fmt.Printf("%2d. %15s (%4d, %2.1f%%) \n", counter, key, val, rel)
		})
	} else {
		out.iterate(20, func(counter int, key string, val int) {
			fmt.Printf("\"%s\",", key)
		})
	}
}

func (input TokenCount) sum() int {
	var total int
	for _, count := range input {
		total += count
	}
	return total
}

func (tokc TokenCount) plot(log bool, pfix string) {
	if len(tokc) > 0 {
		plotObj, err := plot.New()
		if err != nil {
			panic(err)
		}

		plotObj.Title.Text = "Tokens"
		plotObj.X.Label.Text = "Rank"
		plotObj.Y.Label.Text = "Count"

		points := make(plotter.XYs, len(tokc))
		tokc.iterate(len(tokc), func(i int, k string, v int) {
			points[i-1].X = float64(i)
			points[i-1].Y = float64(v)
		})

		err = plotutil.AddScatters(plotObj, "", points)
		if err != nil {
			panic(err)
		}

		fname := "out/tc-" + pfix + ".png"
		if log {
			// plotObj.X.Scale = plot.LogScale{}
			plotObj.Y.Scale = plot.LogScale{}
			fname = "out/tc-" + pfix + ".png"
		}

		// Save the plot to a PNG file.
		if err := plotObj.Save(5*vg.Inch, 5*vg.Inch, fname); err != nil {
			panic(err)
		}
	}
}
