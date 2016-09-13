package main

import "testing"

func TestCompute(t *testing.T) {
	useCases := map[operation][3][3]int{
		add:      [3][3]int{{5, 5, 5}, {6, 6, 6}, {7, 7, 7}},
		multiply: [3][3]int{{4, 4, 4}, {8, 8, 8}, {12, 12, 12}},
	}

	m1 := [3][3]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}}
	m2 := [3][3]int{{4, 4, 4}, {4, 4, 4}, {4, 4, 4}}

	for op, expected := range useCases {
		actual := compute(m1, m2, op)

		if actual != expected {
			t.Fatalf("Operation %d, expected %#v GOT %#v", op, expected, actual)
		}
	}
}

func TestPerform(t *testing.T) {
	useCases := map[operation][3]int{
		add:       {5, 4, 9},
		substract: {5, 4, 1},
		multiply:  {5, 4, 20},
		divide:    {10, 2, 5},
	}

	for op, tuple := range useCases {
		if actual := perform(op, tuple[0], tuple[1]); actual != tuple[2] {
			t.Fatalf("Operation %d, expected %d GOT %d", op, actual, tuple[2])
		}
	}
}
