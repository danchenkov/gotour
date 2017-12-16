package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var closureValueA int
	var closureValueB int
	return func() int {
		closureValueA, closureValueB = closureValueB, closureValueA+closureValueB
		if closureValueB == 0 {
			closureValueB = 1
		}
		return closureValueA
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
