package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotate(t *testing.T) {

	var a, b []int
	var k int

	a = []int{1,2,3,4,5,6,7}
	b = []int{5,6,7,1,2,3,4}
	k = 3
	rotate(a, k)
	for idx, num := range a {
		assert.Equal(t, b[idx], num)
	}

	a = []int{-1,-100,3,99}
	b = []int{3,99,-1,-100}
	k = 2
	rotate(a, k)
	for idx, num := range a {
		assert.Equal(t, b[idx], num)
	}
}
