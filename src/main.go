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

// Global variables
var x int = 20

// Global constants
const pi float64 = 3.14159265358979
const hello = "Hello, World!"

func main() {

	// Different methods for defining variables:
	var s1 string = hello + " 1"
	var s2 = hello + "  2"
	s3 := hello + " 3"

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	s1 = "New " + s1

	fmt.Println(s1)

	// Define some numeric variables:
	var i1 int = 10
	var i2 = 20
	i3 := 3.14
	i4, i5 := 40, 6

	fmt.Println(i1)
	fmt.Println(i2)
	fmt.Println(i3 * float64(i2))
	fmt.Println(i4 / i5)

	// Define some boolean variables:
	var b1 bool = true
	var b2 = false
	b3 := b1 || b2

	fmt.Println(b1)
	fmt.Println(!b2)
	fmt.Println(b3)

	// Define a local constant:
	const g = 9.8

	fmt.Println("\nFor Loop:")
	// For loops:
	for i := 0; i < 30; i++ {

		if i == 6 {
			// Override 6th iteration:
			continue
		}

		if i%2 == 0 {
			fmt.Println(i, "is Even")
		} else {
			fmt.Println(i, "is Odd")
		}

		if i == 10 {
			// Break out of the loop at 10th iteration:
			break
		}
	}

	// Another way to define a for loop:
	j := 12
	for j < 40 {
		fmt.Println(j)
		j += 6
	}

	// If/else statements:
	if num := 9; num < 0 {
		fmt.Println("Negative")
	} else if num < 10 {
		fmt.Println("Has 1 digit")
	} else {
		fmt.Println("Has multiple digits")
	}

	// Switch statements:
	i := 2
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("None")
	}

	// switch without an expression is an alternate
	// way to express if/else logic
	t := 12
	switch {
	case t < 10:
		fmt.Println("Less than 10")
	default:
		fmt.Println("Greater than 10")
	}

	// A type switch compares types instead of values.
	// You can use this to discover the type of
	// an "interface" value.

}
