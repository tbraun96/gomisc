package main

import "os"
import "strconv"

func main() {
	// first element of os.Args is the binary. Ensure the array has at least 2 elements
	if len(os.Args) < 2 {
		println("Please provide an input parameter")
		return;
	}

	// parse the input
	input := os.Args[1];
	initial, err := strconv.Atoi(input);
	if err != nil {
		println("Invalid input. Please supply a positive integer only");
		return;
	}

	// ensure the input is valid for the domain of the fibonacci sequence.
	// Furthermore, we are typecasting the int from atoi to a uint, so only
	// positive values are valid
	if initial < 0 {
		println("Positive integers only")
		return;
	}

	println("Finding fib(", initial, ")");
	// input is valid
	println("Val: ", fib(uint(initial)))
}

// This recursion creates a tree whereby the final result is the sum of the
// count of the child nodes generated. The range of this recursive function
// is isomorphic to the fibonacci sequence for all positive-value integers
func fib(val uint) uint {
	if val <= 1 {
		return val
	}

	return fib(val-1) + fib(val-2)
}