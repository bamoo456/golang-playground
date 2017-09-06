package main

import (
	"testing"
)

type calcData struct {
	data []float64
	want float64
}

func TestCalcAvg(t *testing.T) {
	tests := []struct {
		data []float64
		want float64
	}{
		{[]float64{1, 2, 3}, 2},
		{[]float64{1, 1, 1}, 1},
		{[]float64{1.5, 1.5, 1.5}, 1.5},
		// failure case
		//{[]float64{1.5, 1.5, 1.5}, 2},
	}

	for _, test := range tests {
		if actual := CalcAvg(test.data...); actual != test.want {
			t.Errorf("CalcAvg(): input %v, got actual= %v != want=%v", test.data, actual, test.want)
		}
	}
}

func BenchmarkCalcAvg(b *testing.B) {
	test := struct {
		data []float64
		want float64
	}{
		[]float64{301020.495, 19284.122, 45020101.111},
		2,
	}
	for i := 0; i < b.N; i++ {
		CalcAvg(test.data...)
	}
}
