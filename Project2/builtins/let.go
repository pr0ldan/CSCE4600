// builtins/let.go
package builtins

import (
	"fmt"
	"io"
	"strings"

	"github.com/Knetic/govaluate"
)

// Let evaluates arithmetic expressions and prints the result.
func Let(w io.Writer, args ...string) error {
	if len(args) == 0 {
		fmt.Fprintln(w, "Usage: let expression [expression ...]")
		return nil
	}

	// Concatenate the expressions into a single string.
	combinedExpr := strings.Join(args, " ")

	// Evaluate the combined expression using govaluate.
	expression, err := govaluate.NewEvaluableExpression(combinedExpr)
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
		return err
	}

	result, err := expression.Evaluate(nil)
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
		return err
	}

	fmt.Fprintf(w, "%v\n", result)
	return nil
}
