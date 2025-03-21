package main

import (
	"flag"
	"fmt"
	lab2 "github.com/rmerezha/kpi3-lab2"
	"io"
	"os"
	"strings"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "File with expression to compute")
	outputFile      = flag.String("o", "", "File to write the result to")
)

func main() {
	flag.Parse()

	var input io.Reader
	var output io.Writer = os.Stdout

	if *inputExpression != "" && *inputFile != "" {
		fmt.Fprintln(os.Stderr, "Error: Both -e and -f flags cannot be specified at the same time.")
		os.Exit(1)
	}

	if *inputExpression != "" {
		input = strings.NewReader(*inputExpression)
	} else if *inputFile != "" {
		file, err := os.Open(*inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening input file: %s\n", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	} else {
		fmt.Fprintln(os.Stderr, "Error: Either -e or -f flag must be specified.")
		os.Exit(1)
	}

	if *outputFile != "" {
		file, err := os.Create(*outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file: %s\n", err)
			os.Exit(1)
		}
		defer file.Close()
		output = file
	}

	handler := &lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during computation: %s\n", err)
		os.Exit(1)
	}
}