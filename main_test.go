package main

import (
	"math"
	"testing"
)

func TestCalculateFinancialRate(t *testing.T) {
	testCases := []struct {
		name      string
		op        string
		r         float64
		n         float64
		expected  float64
		expectErr bool
	}{
		{
			name:      "Future Value Factor",
			op:        "Future Value Factor",
			r:         0.05,
			n:         10,
			expected:  math.Pow(1+0.05, 10),
			expectErr: false,
		},
		{
			name:      "Present Value Factor",
			op:        "Present Value Factor",
			r:         0.05,
			n:         10,
			expected:  math.Pow(1+0.05, -10),
			expectErr: false,
		},
		{
			name:      "Future Value of Annuity Factor",
			op:        "Future Value of Annuity Factor",
			r:         0.05,
			n:         10,
			expected:  (math.Pow(1+0.05, 10) - 1) / 0.05,
			expectErr: false,
		},
		{
			name:      "Present Value of Annuity Factor",
			op:        "Present Value of Annuity Factor",
			r:         0.05,
			n:         10,
			expected:  (1 - math.Pow(1+0.05, -10)) / 0.05,
			expectErr: false,
		},
		{
			name:      "Sinking Fund Factor",
			op:        "Sinking Fund Factor",
			r:         0.05,
			n:         10,
			expected:  0.05 / (math.Pow(1+0.05, 10) - 1),
			expectErr: false,
		},
		{
			name:      "Capital Recovery Factor",
			op:        "Capital Recovery Factor",
			r:         0.05,
			n:         10,
			expected:  (0.05 * math.Pow(1+0.05, 10)) / (math.Pow(1+0.05, 10) - 1),
			expectErr: false,
		},
		{
			name:      "Unknown Operation",
			op:        "Unknown",
			r:         0.05,
			n:         10,
			expected:  0,
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("Input: op=%s, r=%f, n=%f", tc.op, tc.r, tc.n)
			rate, err := calculateFinancialRate(tc.op, tc.r, tc.n)
			t.Logf("Output: rate=%f, err=%v", rate, err)

			if (err != nil) != tc.expectErr {
				t.Errorf("expected error: %v, got: %v", tc.expectErr, err)
			}

			if !tc.expectErr && math.Abs(rate-tc.expected) > 1e-9 {
				t.Errorf("expected: %f, got: %f", tc.expected, rate)
			}
		})
	}
}
