package regression

import (
	"math"

	"github.com/Julien2313/GoModels/model"
)

type OlsPoint struct {
	Point                            *model.Point
	ErrorX, ErrorX2, ErrorY, ErrorXY float64
}

type Ols struct {
	MeanX, MeanY float64
	OlsPoints    []OlsPoint
	B1, B2       float64
}

func (o Ols) ConvertPoints(points *[]model.Point) {
	o.OlsPoints = make([]OlsPoint, len(*points))
	for i, _ := range *points {
		o.OlsPoints[i].Point = &(*points)[i]
	}
}

func (o Ols) ComputeMeans() {
	totX := 0.0
	totY := 0.0
	for _, olsPoint := range o.OlsPoints {
		totX += olsPoint.Point.X
		totY += olsPoint.Point.Y
	}
	o.MeanX = totX / float64(len(o.OlsPoints))
	o.MeanY = totY / float64(len(o.OlsPoints))
}

func (o Ols) ComputeErrors() {
	for i, _ := range o.OlsPoints {
		olsPoint := o.OlsPoints[i]

		olsPoint.ErrorX = olsPoint.Point.X - o.MeanX
		olsPoint.ErrorX2 = math.Pow(olsPoint.ErrorX, 2)
		olsPoint.ErrorY = olsPoint.Point.Y - o.MeanY
		olsPoint.ErrorXY = olsPoint.ErrorX * olsPoint.ErrorY
	}
}

func (o Ols) ComputeBs() {
	totErrorX2 := 0.0
	totErrorXY := 0.0
	for i, _ := range o.OlsPoints {
		olsPoint := o.OlsPoints[i]

		totErrorX2 += olsPoint.ErrorX2
		totErrorXY += olsPoint.ErrorXY
	}
	o.B1 = totErrorXY / totErrorX2
}
