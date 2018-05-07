package main

import (
	"fmt"
)

type Matrix struct {
	rows, columns int
	elements      []float64
}

func newMatrix(rows int, columns int, elements []float64) *Matrix {
	return &Matrix{
		rows:     rows,
		columns:  columns,
		elements: elements,
	}
}

func (m Matrix) Transpose() *Matrix {
	newElements := make([]float64, len(m.elements))
	for i := 0; i < m.columns; i++ {
		for j := 0; j < m.rows; j++ {
			newElements[i*m.rows+j] = m.elements[j*m.columns+i]
		}
	}

	return &Matrix{
		rows:     m.columns,
		columns:  m.rows,
		elements: newElements,
	}
}

func (m Matrix) Print() {
	for i := 0; i < m.rows; i++ {
		fmt.Println(m.elements[i*m.columns : (i+1)*m.columns])
	}
}

func (m Matrix) Multiply(n Matrix) *Matrix {
	if m.columns != n.rows {
		return
	}

	newElements := make([]float64, len(m.columns*n.rows))

	return newMatrix()
}
