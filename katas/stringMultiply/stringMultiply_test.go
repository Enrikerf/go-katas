package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringMultipy(testing *testing.T) {
	var valuesToTest = []struct {
		x, y, result string
	}{
		{"12341234", "1234", "15229082756"},
		{"1234", "12341234", "15229082756"},
		{"0", "12341234", "0"},
		{"123123", "0", "0"},
	}
	for i := 0; i < len(valuesToTest); i++ {
		assert.Equal(testing, valuesToTest[i].result, stringMultiply(valuesToTest[i].x, valuesToTest[i].y), "should be equal")
	}

}
