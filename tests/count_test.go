package tests

import (
	"cc/wc-tool/src"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByteCount(t *testing.T) {
	const expected = 3735
	values := src.ReadFile(testFileName)
	actual := src.CountBytes(values)
	assert.Equal(t, expected, actual)

}
func TestLineCount(t *testing.T) {
	const expected = 9
	values := src.ReadFile(testFileName)
	actual := src.CountLines(values)
	assert.Equal(t, expected, actual)
}
func TestWordCount(t *testing.T) {
	const expected = 551
	values := src.ReadFile(testFileName)
	actual := src.CountWords(values)
	assert.Equal(t, expected, actual)
}
func TestCharCount(t *testing.T) {
	const expected = 3735
	values := src.ReadFile(testFileName)
	actual := src.CountChars(values)
	assert.Equal(t, expected, actual)
}
