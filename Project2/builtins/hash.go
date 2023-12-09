package builtins

import (
	"fmt"
)

// HashCommand executes the "hash" built-in command.
func HashCommand(args ...string) error {
	// Parse the options and arguments for the "hash" command.
	// Implement the desired behavior based on the options and arguments.

	// For simplicity, let's print the parsed options and arguments.
	fmt.Printf("Options: %+v\n", parseHashOptions(args))
	fmt.Printf("Arguments: %+v\n", parseHashArguments(args))

	return nil
}

// parseHashOptions parses the options for the "hash" command.
func parseHashOptions(args []string) map[string]interface{} {
	options := make(map[string]interface{})

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-r", "-d", "-t":
			options[args[i]] = true
		case "-p":
			options["-p"] = true
			// If -p is specified, the filename is the next argument.
			if i+1 < len(args) {
				options["filename"] = args[i+1]
				i++ // Skip the next argument since it is the filename.
			}
		}
	}

	return options
}

// parseHashArguments parses the non-option arguments for the "hash" command.
func parseHashArguments(args []string) []string {
	var arguments []string

	for _, arg := range args {
		if arg[0] != '-' {
			arguments = append(arguments, arg)
		}
	}

	return arguments
}
