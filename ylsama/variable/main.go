package main

import (
	"fmt"
	"math/cmplx"
)

// example of var block placement can be outsite main func scope
var (
	toBe   bool       = false
	maxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	// example of var block placement can be inside main func scope
	var (
		explicit string = "Hello, I'm a explicitly declared variable"
		inferred        = "Hello, I'm a variable"
	)
	fmt.Printf("Type: %T Value: %v\n", toBe, toBe)
	fmt.Printf("Type: %T Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Printf("Type: %T Value: %v\n", explicit, explicit)
	fmt.Printf("Type: %T Value: %v\n", inferred, inferred)
	localFunc()
}

func localFunc() {
	// example of var block placement can be inside main func scope
	var (
		inferred = "Hello"
	)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	// This code show undeclare
	// fmt.Printf("Type: %T Value: %v\n", explicit, explicit)
	fmt.Printf("Type: %T Value: %v\n", inferred, inferred)
}
