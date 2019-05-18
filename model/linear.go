package model

type Linear2D struct {
	NbrPoints  int
	MinX, MaxX float64
	//A.x+B
	A, B   float64
	Noise  *Noise
	Points []Point
}
