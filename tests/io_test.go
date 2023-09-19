package tests

import (
	"bytes"
	"cc/wc-tool/src"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	values := src.ReadFile(testFileName)
	assert.Equal(t, len(values), 3735)
}

func TestReadFromStdin(t *testing.T) {
	input := src.ReadFile(testFileName)

	r, w, _ := os.Pipe()
	w.Write(input)
	w.Close()
	// To restore Stdin after the test
	defer func(v *os.File) { os.Stdin = v }(os.Stdin)
	os.Stdin = r

	result := src.ReadFromStdin()
	if !bytes.Equal(result, input) {
		t.Errorf("ReadFromStdin() returned %v, expected %v", result, input)
	}
}
