package tools

import (
	"strings"
)

// simple string builder - mmmmm tasty
// Returns the filler (string) surrounded by the bread (another string)
// Multiple kinds of bread can be used for the same filler
func StringSandwich(filler string, bread... string) []string {
	sandwiches := make([]string, 0)
	for _, b := range bread {
		var sg strings.Builder
		sg.WriteString(b)
		sg.WriteString(filler)
		sg.WriteString(b)
		sandwiches = append(sandwiches, sg.String())
	}
	return sandwiches
}