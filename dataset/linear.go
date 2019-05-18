package dataset

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/Julien2313/GoModels/model"
)

func GenerateRandom2DLinearDataSet(dataSet *model.Linear2D) error {
	if dataSet.NbrPoints <= 0 {
		return errors.New("bad number of points :" + fmt.Sprintf("%d", dataSet.NbrPoints))
	}
	if dataSet.MinX >= dataSet.MaxX {
		return errors.New("bad min/max X :" + fmt.Sprintf("%f/%f", dataSet.MinX, dataSet.MaxX))
	}
	if dataSet.Noise.Variance < 0 {
		return errors.New("bad noise :" + fmt.Sprintf("%f", dataSet.Noise))
	}

	dataSet.Points = make([]model.Point, dataSet.NbrPoints)

	for numPoint := 0; numPoint < dataSet.NbrPoints; numPoint++ {
		x := rand.Float64()*(dataSet.MaxX-dataSet.MinX) + dataSet.MinX
		y := dataSet.A*x + dataSet.B
		y = y + rand.NormFloat64()*dataSet.Noise.Variance + dataSet.Noise.Mean
		dataSet.Points[numPoint] = model.Point{X: x, Y: y}
	}

	return nil
}
