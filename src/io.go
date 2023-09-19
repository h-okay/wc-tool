package src

import (
	"io"
	"os"
)

func ReadFile(filename string) []byte {
	file, _ := os.ReadFile(filename)
	return file
}

func ReadFromStdin() []byte {
	stdin, _ := io.ReadAll(os.Stdin)
	return stdin
}

func GetInput(filename string) []byte {
	if filename != "" {
		return ReadFile(filename)
	}
	return ReadFromStdin()
}
