package repl_test

import (
	"slices"
	"testing"

	"github.com/zoumas/pokecli/internal/repl"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		Description string
		Given       string
		Expected    []string
	}{
		{
			Description: "Split the fields",
			Given:       "hello world",
			Expected:    []string{"hello", "world"},
		},
		{
			Description: "Trim leading and trailing whitespace",
			Given:       " hello  world\n",
			Expected:    []string{"hello", "world"},
		},
		{
			Description: "Lowercase only the first field",
			Given:       " HeLLo   World\n",
			Expected:    []string{"hello", "World"},
		},
	}

	for _, cs := range cases {
		t.Run(cs.Description, func(t *testing.T) {
			actual := repl.CleanInput(cs.Given)
			if !slices.Equal(actual, cs.Expected) {
				t.Errorf("\nexpected: %v\nactual: %v", cs.Expected, actual)
			}
		})
	}
}
