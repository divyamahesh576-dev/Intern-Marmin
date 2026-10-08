package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMul(t *testing.T) {

	tests := []struct {
		x   int
		res int
	}{
		{11, 2},
		{30, 6},
		{60, 12},
		{-15, 0},
		{5, 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {

			actual := multiple(tt.x)
			assert.Equal(t, tt.res, actual)
		})
	}
}
