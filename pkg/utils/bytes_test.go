package utils

import "testing"

func TestConvertBytes(t *testing.T) {
	tests := []struct {
		name  string
		size  int64
		want  float64
		want1 MemoryUnit
	}{
		{
			"B",
			128,
			128,
			"B",
		},
		{
			"Kib",
			1024,
			1,
			"KiB",
		},
		{
			"MiB",
			11048576,
			10.5,
			"MiB",
		},
		{
			"GiB",
			11048257669,
			10.3,
			"GiB",
		},
		{
			"TiB",
			18904825987669,
			17.2,
			"TiB",
		},
		{
			"PiB",
			18904345825987669,
			16.8,
			"PiB",
		},
		{
			"EiB",
			5890434582534987669,
			5.1,
			"EiB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ConvertBytes(tt.size)
			if got != tt.want {
				t.Errorf("ConvertSize() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ConvertSize() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
