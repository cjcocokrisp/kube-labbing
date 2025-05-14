package main

import (
	"fmt"
	"testing"
)

func TestFac(t *testing.T) {
	testCases := []struct {
		input  int
		output int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{6, 720},
		{7, 5040},
		{8, 40320},
		{9, 362880},
		{10, 3628800},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Fib %d", testCase), func(t *testing.T) {
			result := Fac(testCase.input)
			if result != testCase.output {
				t.Errorf("Incorrect result: should be %d, but got %d", testCase.output, result)
			}
		})
	}
}
