package main

type square struct {
	sizeLength float64
}

func (sq square) getArea() float64 {
	return sq.sizeLength * sq.sizeLength
}