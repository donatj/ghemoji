package main

import (
	"flag"
	"os"
	"testing"
)

func setupArgs(t *testing.T, args ...string) {
	t.Helper()

	// store the current arguments
	oldArgs := os.Args

	os.Args = append([]string{"cmd"}, args...)
	flag.Parse()

	// restore the old arguments when the test is over
	t.Cleanup(func() {
		os.Args = oldArgs
	})
}

func TestReadInput(t *testing.T) {
	t.Run("read from a named pipe", func(t *testing.T) {
		// copy named pipe input to stdin
		in := os.Stdin
		// make sure to restore it when done with the test
		// without this, concurrent tests will fail
		t.Cleanup(func() {
			os.Stdin = in
		})

		// this is a way to simulate a named pipe
		r, w, _ := os.Pipe()
		t.Cleanup(func() {
			r.Close()
			w.Close()
		})

		input := "Testing named pipe input\nfoo bar\nbaz"
		w.WriteString(input)
		w.Close() // this one is important as it triggers the EOF sent to stdin

		// now we can set the named pipe as if it was stdin
		// please note it would be restored by the cleanup
		os.Stdin = r

		result, err := readInput()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result != input {
			t.Errorf("Expected: %s, Got: %s", input, result)
		}
	})

	t.Run("read from command-line arguments", func(t *testing.T) {
		for name, c := range map[string]struct {
			Arguments []string
			Expected  string
		}{
			"no arguments": {
				Arguments: []string{},
				Expected:  "",
			},
			"one argument": {
				Arguments: []string{"foo"},
				Expected:  "foo",
			},
			"one argument with multiple words": {
				Arguments: []string{"foo bar baz qux"},
				Expected:  "foo bar baz qux",
			},
			"multiple arguments": {
				Arguments: []string{"foo", "bar", "baz", "qux"},
				Expected:  "foo bar baz qux",
			},
		} {
			t.Run(name, func(t *testing.T) {
				setupArgs(t, c.Arguments...)

				result, err := readInput()
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result != c.Expected {
					t.Errorf("Expected: %s, Got: %s", c.Expected, result)
				}
			})
		}
	})
}
