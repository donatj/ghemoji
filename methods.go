package ghemoji

import "strings"

//go:generate go run cmd/ghupdate/main.go -- emoji.go
//go:generate go fmt emoji.go

// Replace replaces all availible github emoji :string:'s with actual emoji characters
//
// @todo: optimization
func Replace(input string) string {
	for k, r := range emoji {
		input = strings.Replace(input, ":"+k+":", r, -1)
	}

	return input
}
