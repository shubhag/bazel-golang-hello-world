package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		first    int
		second   int
		response int
	}{
		{
			name:     "positive number test case",
			first:    1,
			second:   2,
			response: 3,
		},
		{
			name:     "negative number test case",
			first:    -1,
			second:   -2,
			response: -3,
		},
	}

	for _, tc := range tests {
		got := sum(tc.first, tc.second)
		if tc.response != got {
			t.Fatalf("%s: expected response: %v, got: %v", tc.name, tc.response, got)
		}
	}
}
