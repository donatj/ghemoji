package ghemoji

import "testing"

func TestReplace(t *testing.T) {
	var tests = []struct {
		input, output string
	}{
		{"I approve :+1:", "I approve 👍"},
		{"I approve :+1:+1:", "I approve 👍+1:"},
		{"I approve :-1: :+1: :8ball:", "I approve 👎 👍 🎱"},
	}

	for _, test := range tests {
		actual := ReplaceAll(test.input)
		if actual != test.output {
			t.Errorf(`ReplaceAll(%q) = %q; want %q`, test.input, actual, test.output)
		}
	}
}
