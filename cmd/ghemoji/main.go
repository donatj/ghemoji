package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/donatj/ghemoji"
	"github.com/mattn/go-isatty"
)

func main() {
	flag.Usage = usage
	flag.Parse()

	data := strings.Join(flag.Args(), " ")
	if data != "" {
		fmt.Println(ghemoji.ReplaceAll(data))
		return
	}

	if isatty.IsTerminal(os.Stdin.Fd()) {
		fmt.Fprintf(os.Stderr, "Reading from standard input... (press Ctrl-D to end)\n")
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(ghemoji.ReplaceAll(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
		flag.Usage()
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [text]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  Reads from standard input if no text arguments are provided.\n")
	flag.PrintDefaults()
}
