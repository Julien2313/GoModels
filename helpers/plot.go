package helpers

import (
	"sort"

	"github.com/Julien2313/GoModels/model"
	"gonum.org/v1/plot/plotter"
)

func ConvertoPointForPlotting(pts *[]model.Point) plotter.XYs {
	ptsConverted := make(plotter.XYs, len(*pts))
	sort.Slice(*pts, func(i, j int) bool {
		return (*pts)[i].X < (*pts)[j].X
	})
	for i := range *pts {
		pt := (*pts)[i]
		ptsConverted[i].X = pt.X
		ptsConverted[i].Y = pt.Y
	}
	return ptsConverted
}
