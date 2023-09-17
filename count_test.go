package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testFileName = "test.txt"

func TestByteCount(t *testing.T) {
	const expected = 335039
	actual, _ := CountBytes(testFileName)
	assert.Equal(t, expected, actual)

}
func TestLineCount(t *testing.T) {
	const expected = 7142
	actual, _ := CountLines(testFileName)
	assert.Equal(t, expected, actual)
}
func TestWordCount(t *testing.T) {
	const expected = 58164
	actual, _ := CountWords(testFileName)
	assert.Equal(t, expected, actual)
}
func TestCharCount(t *testing.T) {
	const expected = 332143
	actual, _ := CountChars(testFileName)
	assert.Equal(t, expected, actual)
}
