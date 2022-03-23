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
	arrayVariable()
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

func arrayVariable() {
	const (
		LENGTH = 8
	)
	var (
		a [6]int8 = [6]int8{1, 2, 3, 4, 5, 6} //array
	)
	b := [LENGTH]int8{1, 2}
	fmt.Printf("Type: %T Value: %v Length: %v\n", a, a, len(a))
	fmt.Printf("Type: %T %T %T Value: %v %v %v\n", b, b[0], b[1], b, b[0], b[1])
}
