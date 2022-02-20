// The first statement in a Go source file must be
// package name.
// Executable commands must always use package main.
package main

/*
 * The import keyword is how we include code from
 * other packages to use with our program.
 * The fmt package (shorthand for format) implements
 * formatting for input and output.
 */
import "fmt"

func main() {

	// Different methods for defining variables:
	var s1 string = "Hello, World! 1"
	var s2 = "Hello, World! 2"
	s3 := "Hello, World! 3"

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	s1 = "New " + s1

	fmt.Println(s1)

}
