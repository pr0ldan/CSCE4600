package builtins

import (
	"fmt"
	"io"
	"strconv"
)

// Let evaluates arithmetic expressions and prints the result.
func Let(w io.Writer, args ...string) error {
	if len(args) == 0 {
		fmt.Fprintln(w, "Usage: let expression [expression ...]")
		return nil
	}

	result, err := evaluateExpressions(args...)
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
		return err
	}

	fmt.Fprintf(w, "%d\n", result)
	return nil
}

func evaluateExpressions(expressions ...string) (int, error) {
	// Concatenate the expressions into a single string.
	combinedExpr := ""
	for _, expr := range expressions {
		combinedExpr += expr + " "
	}

	// Evaluate the combined expression using the strconv package.
	result, err := strconv.Atoi(combinedExpr)
	if err != nil {
		return 0, fmt.Errorf("invalid expression: %s", combinedExpr)
	}

	return result, nil
}
