package main

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	testCases := []struct {
		input  int
		output int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
		{11, 89},
		{12, 144},
		{13, 233},
		{14, 377},
		{15, 610},
		{16, 987},
		{17, 1597},
		{18, 2584},
		{19, 4181},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Fib %d", testCase), func(t *testing.T) {
			result := Fib(testCase.input)
			if result != testCase.output {
				t.Errorf("Incorrect result: should be %d, but got %d", testCase.output, result)
			}
		})
	}
}
