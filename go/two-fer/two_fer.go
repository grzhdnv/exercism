// Package twofer is for producing Two for One response with a given name
package twofer

import "fmt"

// ShareWith returns a string response with the given name included.
func ShareWith(name string) string {

	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
