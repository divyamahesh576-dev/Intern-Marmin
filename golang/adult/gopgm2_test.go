package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	actual := adult(10)

	assert.Equal(t, false, actual)
}
