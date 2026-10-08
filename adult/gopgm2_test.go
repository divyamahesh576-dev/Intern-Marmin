package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_is_10(t *testing.T) {
	actual := adult(10)

	assert.Equal(t, false, actual)
}
func Test_is_20(t *testing.T) {
	actual := adult(20)

	assert.Equal(t, true, actual)
}
func Test_is_negative_20(t *testing.T) {
	actual := adult(-20)

	assert.Equal(t, false, actual)
}
func Test_is_0(t *testing.T) {
	actual := adult(0)

	assert.Equal(t, false, actual)