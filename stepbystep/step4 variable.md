basic type
bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte
float32 float64
complex64 complex128 (set of all complex numbers with float32/ float64 real and imaginary parts eg: 2.0i.)

```go
package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
```
array 
slice type

objects that are composed of another objects or types.
pointers 

functions 
define functions as variables and pass them to other functions

interface 