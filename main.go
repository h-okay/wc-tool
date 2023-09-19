package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"cc/wc-tool/src"
)

func main() {
	app := &cli.App{
		Name:           "wc-tool",
		Usage:          "count bytes in a file",
		Description:    "A command-line tool for counting various metrics in a file.",
		UsageText:      "wc-tool [OPTIONS] FILE",
		ArgsUsage:      "FILE",
		DefaultCommand: "",
		Action:         countAction,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:               "bytes",
				Aliases:            []string{"c"},
				Usage:              "Print the byte count",
				DisableDefaultText: true,
			},
			&cli.BoolFlag{
				Name:               "lines",
				Aliases:            []string{"l"},
				Usage:              "Print the line count",
				DisableDefaultText: true,
			},
			&cli.BoolFlag{
				Name:               "words",
				Aliases:            []string{"w"},
				Usage:              "Print the words count",
				DisableDefaultText: true,
			},
			&cli.BoolFlag{
				Name:               "chars",
				Aliases:            []string{"m"},
				Usage:              "Print the chars count",
				DisableDefaultText: true,
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

func countAction(c *cli.Context) error {
	filename := c.Args().First()
	input := src.GetInput(filename)
	activated := parseFlags(c.FlagNames(), c)
	counts := src.GetCounts(activated, input)
	output := src.GenerateOutput(filename, counts["lines"], counts["words"], counts["chars"], counts["bytes"])
	fmt.Println(output)
	return nil
}

func parseFlags(givenFlags []string, c *cli.Context) map[string]bool {
	availableFlagNames := []string{"bytes", "lines", "words", "chars"}
	activated := make(map[string]bool, 4)

	// No flags given, all activated
	if len(givenFlags) == 0 {
		for _, flagName := range availableFlagNames {
			activated[flagName] = true
		}
	}

	// Only given flags activated
	if len(givenFlags) > 0 {
		for _, flag := range givenFlags {
			activated[flag] = true
		}
	}

	return activated
}
