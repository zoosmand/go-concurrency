package math

import "testing"

type testpair struct {
	values  []float64
	average float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1.0, 1.1, 1.2, 1.6, 1.7}, 1.32},
	{[]float64{-1, 1, 2, -2}, 0.1},
}

func TestAverage(t *testing.T) {
	// v := Average([]float64{1, 2})
	// if v != 1.5 {
	// 	t.Error("Expected 1.5, got", v)
	// }

	if len(tests) == 0 {
		t.Fatal("Testing list is empty")
	}

	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

// gopls -rpc.trace -v check math_test.go
