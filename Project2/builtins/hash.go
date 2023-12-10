package builtins

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"sync"
)

var (
	hashTable = make(map[string]string)
	mutex     sync.Mutex
)

// HashCommand executes the "hash" built-in command.
func HashCommand(args ...string) error {
	// Parse the options and arguments for the "hash" command.
	// Implement the desired behavior based on the options and arguments.
	options := parseHashOptions(args)
	arguments := parseHashArguments(args)

	if val, exists := options["-p"]; exists && val == true {
		if len(arguments) == 0 {
			return fmt.Errorf("filename argument is required with -p option")
		}
		filename := arguments[0]
		if err := loadHashesFromFile(filename); err != nil {
			return err
		}
		return nil
	}

	// Print the current hash table if -l is specified.
	if val, exists := options["-l"]; exists && val == true {
		printHashTable()
		return nil
	}

	// Process name arguments and update the hash table.
	for _, name := range arguments {
		// If -r is specified, remove the entry from the hash table.
		if val, exists := options["-r"]; exists && val == true {
			removeHashEntry(name)
		} else {
			// Search for the full path of the command and update the hash table.
			fullPath, err := searchCommand(name)
			if err != nil {
				return err
			}
			updateHash(name, fullPath)
		}
	}

	return nil
}

// parseHashOptions parses the options for the "hash" command.
func parseHashOptions(args []string) map[string]interface{} {
	options := make(map[string]interface{})

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-r", "-d", "-t", "-l":
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

// removeHashEntry removes an entry from the hash table.
func removeHashEntry(name string) {
	mutex.Lock()
	defer mutex.Unlock()

	delete(hashTable, name)
}

// searchCommand searches for the full path of a command in $PATH.
func searchCommand(name string) (string, error) {
	// Check if the command already has a full path in the hash table.
	if path, exists := hashTable[name]; exists {
		return path, nil
	}

	// Search for the command in $PATH.
	path, err := exec.LookPath(name)
	if err != nil {
		return "", fmt.Errorf("command '%s' not found in $PATH", name)
	}

	return path, nil
}

// loadHashesFromFile loads previously remembered pathnames from a file.
func loadHashesFromFile(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading file '%s': %v", filename, err)
	}

	// Parse the data and update the hash table.
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			updateHash(fields[0], fields[1])
		}
	}

	return nil
}

// printHash prints the current mapping for a command.
func printHash(name string) {
	if path, exists := hashTable[name]; exists {
		fmt.Printf("%s: %s\n", name, path)
	} else {
		fmt.Printf("%s not found in hash table\n", name)
	}
}

// updateHash updates the hash table with the full path of a command.
func updateHash(name, fullPath string) {
	hashTable[name] = fullPath
}

// printHashTable prints the current hash table.
func printHashTable() {
	mutex.Lock()
	defer mutex.Unlock()

	fmt.Println("Current Hash Table:")
	for name, path := range hashTable {
		fmt.Printf("%s: %s\n", name, path)
	}
}
