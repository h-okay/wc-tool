package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "wc-tool",
		Usage: "count bytes in a file",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "bytes",
				Aliases: []string{"c"},
				Usage:   "Print the byte count",
			},
			&cli.BoolFlag{
				Name:    "lines",
				Aliases: []string{"l"},
				Usage:   "Print the line count",
			},
			&cli.BoolFlag{
				Name:    "words",
				Aliases: []string{"w"},
				Usage:   "Print the words count",
			},
			&cli.BoolFlag{
				Name:    "chars",
				Aliases: []string{"m"},
				Usage:   "Print the chars count",
			},
		},
		DefaultCommand: "help",
		Action: func(c *cli.Context) error {
			filename := c.Args().First()
			// Check if the --bytes or -c flag is set
			if c.Bool("bytes") {
				countAndPrint("bytes", filename)
			}
			// Check if the --lines or -l flag is set
			if c.Bool("lines") {
				countAndPrint("lines", filename)
			}
			// Check if the --words or -w flag is set
			if c.Bool("words") {
				countAndPrint("words", filename)
			}
			// Check if the --words or -w flag is set
			if c.Bool("chars") {
				countAndPrint("chars", filename)
			}

			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

func countAndPrint(mode string, filename string) {
	switch mode {
	case "bytes":
		bytes, err := CountBytes(filename)
		printCount(filename, "bytes", bytes, err)
	case "lines":
		lines, err := CountLines(filename)
		printCount(filename, "lines", lines, err)
	case "words":
		words, err := CountWords(filename)
		printCount(filename, "words", words, err)
	case "chars":
		chars, err := CountChars(filename)
		printCount(filename, "chars", chars, err)
	}
}

func printCount(filename string, mode string, count int, err error) {
	if err != nil {
		fmt.Printf("Error counting words for file %s", filename)
		panic(err)
	}
	fmt.Printf("%d %s\n", count, filename)
}
