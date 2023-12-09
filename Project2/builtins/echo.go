package builtins

import (
	"fmt"
	"io"
)

// Echo prints the arguments to the standard output.
func Echo(w io.Writer, args ...string) error {
	// Convert the []string slice to []interface{} for fmt.Fprintln.
	var interfaces []interface{}
	for _, arg := range args {
		interfaces = append(interfaces, arg)
	}

	// Use fmt.Fprintln with the converted slice.
	fmt.Fprintln(w, interfaces...)

	return nil
}
