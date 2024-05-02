package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/donatj/ghemoji"
)

func main() {
	// right now, we are not using flag, but we could
	// so let's plan it
	flag.Usage = usage
	flag.Parse()

	input, err := readInput()
	if err != nil {
		flag.Usage()
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	fmt.Fprintln(os.Stdout, ghemoji.ReplaceAll(input))
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [text]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "or pipe text to the program\n")
	flag.PrintDefaults()
}

// readInput reads input from stdin if stdin is a named pipe.
// Otherwise, it joins the command-line arguments with a space separator.
func readInput() (string, error) {
	info, err := os.Stdin.Stat()
	if err != nil {
		return "", fmt.Errorf("getting stat from stdin: %w", err)
	}

	if info.Mode()&os.ModeNamedPipe != 0 {
		input, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return "", fmt.Errorf("reading from stdin: %wv", err)
		}
		return string(input), nil
	}

	return strings.Join(flag.Args(), " "), nil
}
