package main

import (
	"fmt"
	"os"

	"github.com/gawakawa/cert-from-scratch/basecert"
	"github.com/gawakawa/cert-from-scratch/util"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "basecert":
		if len(os.Args) < 3 {
			fmt.Fprintf(os.Stderr, "Error: output path required\n")
			fmt.Fprintf(
				os.Stderr,
				"Usage: %s basecert <output-path>\n",
				os.Args[0],
			)
			os.Exit(1)
		}
		cert := basecert.New()
		if err := util.MarshalAndSaveCert(os.Args[2], cert); err != nil {
			fmt.Fprintf(
				os.Stderr,
				"Error saving base certificate: '%s'\n",
				os.Args[1],
			)
			printUsage()
			os.Exit(1)
		}
	default:
		fmt.Fprintf(
			os.Stderr,
			"Error: unknown command '%s'\n",
			os.Args[1],
		)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Fprintf(
		os.Stderr,
		"Usage: %s <command> [args...]\n",
		os.Args[0],
	)
	fmt.Fprintf(os.Stderr, "\nCommands:\n")
	fmt.Fprintf(
		os.Stderr,
		"  basecert <output-path>  Generate base certificate in DER format\n",
	)
}
