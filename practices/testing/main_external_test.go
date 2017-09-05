package main_test

// This is the usage of "main_test" (external test package)

import (
	. "golang-playground/practices/testing"
	"testing"
)

func TestIsString(t *testing.T) {
	var tests = []struct {
		data interface{}
		want bool
	}{
		{"v", true},
		{1, false},
		{[]int{1, 2, 3}, false},
		{"Hello world", true},
		// failure case
		{30.6, true},
	}

	for _, test := range tests {
		if actual := IsString(test.data); actual != test.want {
			t.Errorf("IsString(): input %v, got actual= %v not equal to want=%v", test.data, actual, test.want)
		}
	}
}
