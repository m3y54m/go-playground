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
import (
	"fmt"
	"reflect"
	"unicode/utf8"
)

// Global variables
var x int = 20

// Global constants
const pi float64 = 3.14159265358979
const hello = "Hello, World!"

func main() {

	// Different methods for defining variables:
	fmt.Println("\n* VARIABLES & CONSTANTS *")

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

	// For loops:
	fmt.Println("\n* FOR *")

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
	fmt.Println("\n* IF/ELSE *")

	if num := 9; num < 0 {
		fmt.Println("Negative")
	} else if num < 10 {
		fmt.Println("Has 1 digit")
	} else {
		fmt.Println("Has multiple digits")
	}

	// Switch statements:
	fmt.Println("\n* SWITCH *")

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

	// Arrays:
	fmt.Println("\n* ARRAYS *")

	var a [5]uint8
	fmt.Println(a)

	// Modify the array elements:
	a[4] = 100
	fmt.Println(a)
	fmt.Println(a[4])

	// Length of the array:
	fmt.Println(len(a))

	// Initialize an array with values:
	b := [5]int{0, 10, 20, 30, 40}
	fmt.Println(b)

	// Array slicing:
	c := b[1:3]
	fmt.Println(c)

	// Multi-dimensional arrays
	d := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(d)

	// Slices:
	// Unlike arrays, slices are typed only by the
	// elements they contain (not the number of elements).
	fmt.Println("\n* SLICES *")

	s := make([]int, 5)
	fmt.Println(s)

	// Modify the slice elements:
	s[4] = 100
	fmt.Println(s)
	fmt.Println(s[4])

	// Length of the slice:
	fmt.Println(len(s))

	// Append to the slice:
	s = append(s, 200)
	fmt.Println(s)

	// Slices can be composed of multiple
	// elements:
	s = append(s, 300, 400)
	fmt.Println(s)

	// Slices can be copied:
	new_s := make([]int, len(s))
	copy(new_s, s)
	fmt.Println(new_s)

	s[0] = 1000
	fmt.Println(s)
	fmt.Println(new_s)

	// Slicing a slice:
	fmt.Println(s[2:5])
	fmt.Println(s[2:])
	fmt.Println(s[:5])

	// Initialize a slice with an array:
	e := []int{1, 2, 3, 4, 5}
	fmt.Println(e)
	fmt.Println(len(e))

	// Multi-dimensional slices:
	// The length of the inner slices can vary,
	// unlike with multi-dimensional arrays.
	r := make([][]int, 3)
	for i := 0; i < 3; i++ {
		inner_len := i + 1
		r[i] = make([]int, inner_len)
		for j := 0; j < inner_len; j++ {
			r[i][j] = i + j
		}
	}
	fmt.Println(r)

	y := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(y)

	// Maps:
	fmt.Println("\n* MAPS *")

	// Define a map:
	m := make(map[string]int)
	fmt.Println(m)

	m["key1"] = 1
	m["key2"] = 20
	m["key3"] = 13
	fmt.Println(m)
	fmt.Println(len(m))

	// Delete a key:
	delete(m, "key1")
	fmt.Println(m)
	fmt.Println(len(m))

	value, present := m["key1"]
	fmt.Println(value, present)

	value, present = m["key2"]
	fmt.Println(value, present)

	// Initialize a map with a map literal:
	n := map[string]int{"key1": 1, "key2": 2, "key3": 3}
	fmt.Println(n)

	// Range:
	fmt.Println("\n* RANGE *")

	// Range over a map:
	for key, value := range m {
		fmt.Println(key, value)
	}

	// Range over a map with keys only:
	for key := range m {
		fmt.Println(key)
	}

	// Range over a map with values only:
	for _, value := range m {
		fmt.Println(value)
	}

	// Range over a slice:
	for index, value := range s {
		fmt.Println(index, value)
	}

	// Range over string:
	for index, value := range "سلام دنیا" {
		fmt.Println(index, value, string(value))
	}

	// Functions:
	fmt.Println("\n* FUNCTIONS *")

	// Function without return values:
	TestFunc()

	MyPrint(hello)

	// Functions with return values:
	fmt.Println(plus(1, 2))

	// Variadic Function:
	fmt.Println(MySum(1, 2, 3))
	fmt.Println(MySum(1, 2, 3, 4, 5))
	// Pass a slice as an argument:
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(MySum(nums...))

	// Functions with named return values:
	p, t := MySum2(1, 2, 3, 4, 5)
	fmt.Println(p, t)

	// Closures (anonymous functions):
	// A closure is a function value that references
	// variables from outside its body.
	// Anonymous functions are useful when you want to
	// define a function inline without having to name it.
	fmt.Println("\n* CLOSURES *")

	// Anonymous function:
	f := func(x int) int {
		return x * x
	}
	fmt.Println(f(3))

	// Closure with local static variable i
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInt2 := intSeq()
	fmt.Println(nextInt2())

	f2 := func() func() int {
		i := 0
		return func() int {
			i++
			return i
		}
	}
	fmt.Println(f2()())
	fmt.Println(f2()())

	t2 := f2()
	fmt.Println(t2())
	fmt.Println(t2())

	// Recursion:
	fmt.Println("\n* RECURSION *")

	fmt.Println(factorial(3))

	// Recursion with a closure:
	var fibonacci func(n int) int

	fibonacci = func(n int) int {
		if n < 2 {
			return n
		}
		return fibonacci(n-1) + fibonacci(n-2)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(fibonacci(i))
	}

	// Pointers:
	fmt.Println("\n* POINTERS *")

	// Declare a pointer:
	x1 := 10
	var p1 *int = &x1

	fmt.Println("Address of x1:", p1)
	fmt.Println("Value of x1:", x1)

	*p1 = 20

	fmt.Println("Address of x1:", p1)
	fmt.Println("Value of x1:", x1)

	f3 := func(x *int) {
		*x++
	}

	f3(&x1)
	fmt.Println("Value of x1:", x1)

	// Pointers with slices:
	fmt.Println("\n* POINTERS WITH SLICES *")

	// Declare a slice:
	s5 := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", s1)

	// Declare a pointer to a slice:
	p2 := &s5

	// Declare a pointer to a slice element:
	p3 := &(*p2)[4]
	fmt.Println("Address of p2:", p2)
	fmt.Println("Address of p3:", p3)
	fmt.Println("Value of p3:", *p3)

	// Strings and runes:
	fmt.Println("\n* STRINGS AND RUNES *")

	// Declare a string:
	// In Go, a string is a "read-only" slice of bytes.
	const s9 string = "hello سلام"

	fmt.Println(s9)

	// Length of a string (number of bytes):
	fmt.Println(len(s9))

	// Convert a string to a slice of bytes:
	b9 := []byte(s9)
	for index, value := range b9 {
		fmt.Printf("%d %s %x\n", index, string(value), value)
	}

	// Convert a slice of bytes to a string:
	s10 := string(b9)
	fmt.Println(s10)

	// Number of runes (Uincode characters)):
	fmt.Println(len([]rune(s9)))

	// in Go "rune" is an alias for int32 and
	// represents a Unicode character
	// Runes are UTF-8 encoded

	// Range over runes of a string:
	for index, runeValue := range s9 {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}

	// Go stores the UTF-8 encoded byte sequences of string values in memory.
	// The builtin function len() reports the byte-length of a string,
	// so basically the memory required to store a string value in memory is:
	// sizeOfString := len(s9) + int(unsafe.Sizeof(s9))
	// WARNING: UNSAFE CODE!
	sizeOfString := len(s9) + int(reflect.TypeOf(s9).Size())
	fmt.Println("Memory allocated for storing string data type and value:", sizeOfString)
	fmt.Println(reflect.TypeOf(s9))

	// Strings are immutable:
	// s9[0] = 'a' // ERROR: cannot assign to s9[0]

	// Using utf8 package:
	fmt.Println("\n* UTF-8 PACKAGE *")

	fmt.Println(utf8.RuneCount(b9))
	fmt.Println(utf8.RuneCountInString(s9))
	fmt.Println([]rune(s9))

	for index, value := range b9 {
		fmt.Printf("%d %x\n", index, value)
	}

	for index, runeValue := range s9 {
		fmt.Printf("%#U starts at byte position %d | %d\n", runeValue, index, reflect.TypeOf(runeValue).Size())
	}

	// Creating a rune (int32 for Unicode characters)
	r1 := 'h'
	r2 := 'س'
	r3 := '🐫'
	r4 := '\n'
	fmt.Println(r1, string(r1))
	fmt.Println(r2, string(r2))
	fmt.Println(r3, string(r3))
	fmt.Println(r4, string(r4))
	fmt.Printf("%d %#x %#U\n", r3, r3, r3)
	fmt.Println(reflect.TypeOf(r1).Size())

	b4 := []byte(string(r3))
	fmt.Printf("%x\n", b4)
	for index, value := range b4 {
		fmt.Printf("%d %X %d\n", index, value, value)
	}

	s6 := string(b4)
	fmt.Println(s6)

}

// Define some functions:
func TestFunc() {
	fmt.Println("x =", x)
}

func MyPrint(s string) {
	fmt.Println(s)
}

func plus(a int, b int) int {
	return a + b
}

func MySum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func MySum2(nums ...int) (int, int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum, sum * 2
}

// Closures:
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// Recursive function:
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
