package main

import (
	"strings"
	"unicode/utf8"
)

func CountBytes(filename string) (int, error) {
	file, err := ReadFile(filename)
	if err != nil {
		return 0, err
	}
	return len(file), nil
}

func CountLines(filename string) (int, error) {
	file, err := ReadFile(filename)
	if err != nil {
		return 0, err
	}
	return len(strings.Split(string(file), "\n")) - 1, nil
}

func CountWords(filename string) (int, error) {
	file, err := ReadFile(filename)
	if err != nil {
		return 0, err
	}
	return len(strings.Fields(string(file))), nil
}

func CountChars(filename string) (int, error) {
	file, err := ReadFile(filename)
	if err != nil {
		return 0, err
	}
	return utf8.RuneCountInString(string(file)), nil

}
