package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_is_5(t *testing.T) {
	actual := sub(10, 5)

	assert.Equal(t, 5, actual)
}
func Test_is_negative_17(t *testing.T) {
	actual := sub(-11, -6)

	assert.Equal(t, -17, actual)
}
func Test_is_negative_one(t *testing.T) {
	actual := sub(0, 1)

	assert.Equal(t, -1, actual)
}
