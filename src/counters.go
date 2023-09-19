package src

import (
	"strings"
	"unicode/utf8"
)

func GetCounts(activatedModes map[string]bool, values []byte) map[string]int {
	counts := make(map[string]int)
	for mode, activated := range activatedModes {
		if activated {
			counts[mode] = GetModeResult(mode, values)
		}
	}
	return counts
}

func CountBytes(values []byte) int {
	return len(values)
}

func CountLines(values []byte) int {
	return len(strings.Split(string(values), "\n")) - 1
}

func CountWords(values []byte) int {
	return len(strings.Fields(string(values)))
}

func CountChars(values []byte) int {
	return utf8.RuneCountInString(string(values))

}
