package utils

import "testing"

func TestRoundFloat(t *testing.T) {
	tests := []struct {
		name      string
		val       float64
		precision uint
		want      float64
	}{
		{
			"positive with precision 2",
			12.3456789,
			2,
			12.35,
		},
		{
			"positive with precision 5",
			12.3456789,
			5,
			12.34568,
		},
		{
			"negative with precision 0",
			-12.3456789,
			0,
			-12,
		},
		{
			"negative with precision 10",
			-12.3456789,
			10,
			-12.3456789,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RoundFloat(tt.val, tt.precision); got != tt.want {
				t.Errorf("RoundFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
