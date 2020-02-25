package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {

	a := []int{1, 2}
	assert.Equal(t, 2, removeDuplicates(a))

	a = []int{1, 1, 2}
	assert.Equal(t, 2, removeDuplicates(a))

	a = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	assert.Equal(t, 5, removeDuplicates(a))
}
