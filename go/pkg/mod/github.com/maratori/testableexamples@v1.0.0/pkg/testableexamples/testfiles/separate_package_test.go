package testfiles_test

import "fmt"

func Example_good() {
	fmt.Println("hello")
	// Output: hello
}

func Example_goodEmptyOutput() {
	fmt.Println("")
	// Output:
}

func Example_bad() { // want `^missing output for example, go test can't validate it$`
	fmt.Println("hello")
}
