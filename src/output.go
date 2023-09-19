package src

import "fmt"

func GetModeResult(mode string, values []byte) int {
	switch mode {
	case "bytes":
		return CountBytes(values)
	case "lines":
		return CountLines(values)
	case "words":
		return CountWords(values)
	case "chars":
		return CountChars(values)
	}
	return -1
}

func GenerateOutput(filename string, args ...int) string {
	output := "  "
	for _, arg := range args {
		if arg == 0 {
			continue
		}
		output += fmt.Sprint(arg) + " "
	}
	return output + filename
}
