package main

import "testing"

func TestIsHealthy(t *testing.T) {
	tests := []struct {
		desc     string
		results  []bool
		expected bool
	}{
		{
			desc:     "All true",
			results:  []bool{true, true, true, true, true, true, true, true, true, true},
			expected: true,
		},
		{
			desc:     "One false",
			results:  []bool{true, true, true, true, true, true, true, true, true, false},
			expected: true,
		},
		{
			desc:     "Two false",
			results:  []bool{true, true, true, true, true, true, true, true, false, false},
			expected: true,
		},
		{
			desc:     "Five false",
			results:  []bool{true, true, true, true, true, false, false, false, false, false},
			expected: true,
		},
		{
			desc:     "Six false",
			results:  []bool{true, true, true, true, false, false, false, false, false, false},
			expected: false,
		},
		{
			desc:     "All false",
			results:  []bool{false, false, false, false, false, false, false, false, false, false},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			lastResults = tt.results

			actual := isHealthy()
			if actual != tt.expected {
				t.Errorf("isHealthy() = %v, want %v", actual, tt.expected)
			}
		})
	}
}
