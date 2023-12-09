// builtins/pwd.go
package builtins

import (
	"fmt"
	"io"
	"os"
)

// PrintWorkingDirectory prints the current working directory.
func PrintWorkingDirectory(w io.Writer, args ...string) error {
	// No need to check args, as 'pwd' doesn't take any arguments.

	// Get the current working directory.
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	// Print the current working directory.
	_, err = fmt.Fprintf(w, "%s\n", wd)
	return err
}
