// builtins/help.go
package builtins

import (
	"fmt"
	"io"
)

// Help prints help information for the shell built-in commands.
func Help(w io.Writer, args ...string) error {
	// Provide help information for each built-in command.
	helpText := `
Available built-in commands:
- cd [directory]: Change the current working directory.
- env: Print the environment variables.
- pwd: Print pathname of current working directory.
- times: Print system times used.
- hash: Store pathnames of command arguments.
		hash [-r] [-p filename] [-dt] [name]
		-p inhibits path search
		-r forget all locations
		-d forget remembered location of each name
		-t prints full pathname of corresponding name
		-l print remembered commands
- let: Perform arithmetic operation and print.
- echo: Print statements. 
- exit: Exit the shell.
- help: Show help information.

Usage:
  help [command]

Examples:
  help cd
  help env
  help exit
`
	if len(args) == 0 {
		// If no specific command is provided, print general help.
		fmt.Fprintln(w, helpText)
	} else {
		// If a specific command is provided, print help for that command.
		switch args[0] {
		case "cd":
			fmt.Fprintln(w, "cd: Change the current working directory.")
			fmt.Fprintln(w, "Usage: cd [directory]")
		case "env":
			fmt.Fprintln(w, "env: Print the environment variables.")
			fmt.Fprintln(w, "Usage: env")
		case "pwd":
			fmt.Fprintln(w, "pwd: Print pathname of current working directory.")
			fmt.Fprintln(w, "Usage: pwd")
		case "times":
			fmt.Fprintln(w, "times: Print system times used.")
			fmt.Fprintln(w, "Usage: times")
		case "hash":
			fmt.Fprintln(w, "hash: Store pathnames of command arguments.")
			fmt.Fprintln(w, "Usage: hash [-p filename]")
		case "let":
			fmt.Fprintln(w, "let: Perform arithmetic operation and print.")
			fmt.Fprintln(w, "Usage: let [expression]")
		case "echo":
			fmt.Fprintln(w, "echo: Print statements following echo command")
			fmt.Fprintln(w, "Usage: echo [statement]")
		case "exit":
			fmt.Fprintln(w, "exit: Exit the shell.")
			fmt.Fprintln(w, "Usage: exit")
		case "help":
			fmt.Fprintln(w, "help: Show help information.")
			fmt.Fprintln(w, "Usage: help [command]")
		default:
			fmt.Fprintf(w, "Unknown command: %s\n", args[0])
		}
	}

	return nil
}
