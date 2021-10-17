package main

import (
	"context"
	"github.com/kenesparta/tkGrpcService/gen/multiply"
	"math"
	"testing"
)

func TestResultMultiplyGetsCorrectValue(t *testing.T) {
	s := Server{}

	tests := []struct {
		First  float64
		Second float64
		Output float64
	}{
		{
			First:  2.1,
			Second: 14.2,
			Output: 29.820,
		},
		{
			First:  -27788.001,
			Second: 14138824.278,
			Output: -392889663175.888,
		},
	}

	for _, tt := range tests {
		resp, err := s.Multiply(context.Background(), &multiply.Factor{
			First:  tt.First,
			Second: tt.Second,
		})
		if err != nil {
			t.Errorf("HelloTest(%v) got unexpected error", err)
		}
		if resp.Result != tt.Output {
			t.Errorf("%.3f is different than %.3f", resp.Result, tt.Output)
		}
	}
}

func TestResultMultiplyReturnsInfiniteError(t *testing.T) {
	s := Server{}

	tests := []struct {
		First  float64
		Second float64
	}{
		{
			First:  math.MaxFloat64,
			Second: 14.2,
		},
		{
			First:  -27788.001,
			Second: math.MaxFloat64,
		},
	}

	for _, tt := range tests {
		_, err := s.Multiply(context.Background(), &multiply.Factor{
			First:  tt.First,
			Second: tt.Second,
		})
		if err.Error() != "the product is infinite" {
			t.Errorf("HelloTest(%v) got unexpected error", err)
		}
	}
}
