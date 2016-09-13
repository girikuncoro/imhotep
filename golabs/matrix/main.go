/*
Write a program (with tests!) to compute 3x3 matrix addition and multiplication.
Extra credit:support substract and divide operations
*/
package main

import "fmt"

type operation int

const (
	add operation = iota
	substract
	multiply
	divide
)

func compute(m1, m2 [3][3]int, op operation) (res [3][3]int) {
	for i := 0; i < len(m1); i++ {
		for j := 0; j < len(m1[i]); j++ {
			res[i][j] = perform(op, m1[i][j], m2[i][j])
		}
	}
	return
}

func perform(op operation, a, b int) (res int) {
	switch op {
	case add:
		res = a + b
	case multiply:
		res = a * b
	case substract:
		res = a - b
	case divide:
		res = a / b
	}
	return
}

func main() {
	m1 := [3][3]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}}
	m2 := [3][3]int{{3, 3, 3}, {4, 4, 4}, {5, 5, 5}}

	fmt.Println(compute(m1, m2, add))
	fmt.Println(compute(m1, m2, substract))
	fmt.Println(compute(m1, m2, multiply))
	fmt.Println(compute(m1, m2, divide))
}
