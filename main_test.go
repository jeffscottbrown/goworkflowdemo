package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessage(t *testing.T) {
	result := getMessage()
	expected := "Hello, World!"

	assert.Equal(t, expected, result, "The message should be 'Hello, World!'")
}
