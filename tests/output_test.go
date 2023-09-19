package tests

import (
	"cc/wc-tool/src"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratedOutputNoArgs(t *testing.T) {
	expected := "  " + testFileName
	actual := src.GenerateOutput(testFileName)
	assert.Equal(t, expected, actual)
}

func TestGeneratedOutputWithArgs(t *testing.T) {
	expected := "  123 123 " + testFileName
	actual := src.GenerateOutput(testFileName, 123, 123)
	assert.Equal(t, expected, actual)
}
