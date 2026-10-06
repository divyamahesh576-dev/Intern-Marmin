package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSub(t *testing.T) {
	actual := sub(10, 5)

	assert.Equal(t, 5, actual)
}
