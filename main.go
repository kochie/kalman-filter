package main

func main() {
	a := newMatrix(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	b := a.Transpose()
	b.Print()
	a.Print()
}
