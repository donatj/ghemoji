# ghemoji

[![GoDoc](https://godoc.org/github.com/donatj/ghemoji?status.svg)](https://godoc.org/github.com/donatj/ghemoji)
[![CI](https://github.com/donatj/ghemoji/actions/workflows/ci.yml/badge.svg)](https://github.com/donatj/ghemoji/actions/workflows/ci.yml)

Go library for working with GitHub's `:emoji:`

Please refer to documentation https://pkg.go.dev/github.com/donatj/ghemoji

## CLI Tool

A simple CLI tool is included for easy scripting.

```
go install github.com/donatj/ghemoji/cmd/ghemoji@latest
```

### Example usage

Read from arguments:

```
$ ghemoji ":+1:"
👍
```

Read file standard in:

```
$ ghemoji < example.md
Example text from standard in 😄
```

Read from standard in:

```
$ ghemoji
Reading from standard input... (press Ctrl-D to end)
:wink:  
😉
test :100: test 
test 💯 test
```
