package main

import (
	"math/rand"

	"github.com/Julien2313/GoModels/dataset"
	"github.com/Julien2313/GoModels/helpers"
	"github.com/Julien2313/GoModels/model"
	"github.com/Julien2313/GoModels/model/regression"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	rand.Seed(helpers.GenerateSeed())
	dataSet := &model.Linear2D{
		NbrPoints: 100,
		MinX:      -10,
		MaxX:      10,
		A:         rand.Float64()*2.0 - 1.0,
		B:         rand.Float64()*2.0 - 1.0,
		Noise: &model.Noise{
			Mean:     0,
			Variance: 1,
		},
	}
	err := dataset.GenerateRandom2DLinearDataSet(dataSet)
	if err != nil {
		panic(err)
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	ols := &regression.Ols{}
	ols.Compute(&dataSet.Points, p)

	p.Title.Text = "Regression OLS"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	p.X.Min = -10
	p.X.Max = 10
	p.Y.Min = -20
	p.Y.Max = 20

	err = plotutil.AddLinePoints(p,
		"Data", helpers.ConverPointtoPointForPlotting(&dataSet.Points))
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "out/points.png"); err != nil {
		panic(err)
	}
}
