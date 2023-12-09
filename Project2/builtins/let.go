package builtins

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Let evaluates arithmetic expressions and prints the result.
func Let(w io.Writer, args ...string) error {
	if len(args) == 0 {
		fmt.Fprintln(w, "Usage: let expression [expression ...]")
		return nil
	}

	// Concatenate the expressions into a single string.
	combinedExpr := strings.Join(args, " ")

	// Evaluate the combined expression using the strconv package.
	result, err := strconv.Atoi(combinedExpr)
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
		return err
	}

	fmt.Fprintf(w, "%d\n", result)
	return nil
}
