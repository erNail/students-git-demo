package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	actualResult := CalculateSum(3, 5)
	expectedResult := 8

	assert.Equal(t, expectedResult, actualResult)
}

func TestProduct(t *testing.T) {
	actualResult := CalculateProduct(3, 5)
	expectedResult := 15

	assert.Equal(t, expectedResult, actualResult)
}
