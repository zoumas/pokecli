package repl

import "strings"

// CleanInput trims all leading and trailing whitespace from s using strings.TrimSpace.
// It then splits s into fields using strings.Fields and lowercase only the first field which is to be the command separating it from its arguments.
func CleanInput(s string) []string {
	fs := strings.Fields(strings.TrimSpace(s))
	if len(fs) > 0 {
		fs[0] = strings.ToLower(fs[0])
	}
	return fs
}
