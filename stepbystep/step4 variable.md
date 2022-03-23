# basic type
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
# array 
sequence of elements of a single type, size cannot be changed.
```go
//define
var arr [100]int
//declare
arrayVar1 := [8]int8{0, 1,2,3,4,..}
arrayVar2 [8]int8 := [8]int8{0, 1,2,3,4,..}
//define then declare
var arr [2]bool
arr[0] = true
arr[1] = false
```
## zero-initialization: 
Go will initialize every value: bool to false, int to 0,...

# slice type
arrays, but size can be changed.
```go
mySlice := make([]int, 10)

mySlice := append(mySlice, 5)
mySlice := mySlice[1:]
mySlice = append(mySlice[:1], mySlice[2:]...)
```
# Maps
dictionaries–for each word ```myMap := make(map[<type1|key>]<type2|value>)```  
type cannot use as key: slices, the functions, and maps
```go
	myMap := make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2
	fmt.Println(myMap["one"])
```
## interface map
JSON Unmarshal
```go
	myJsonMap := make(map[string]interface{})
	jsonData := []byte(`{"hello":"world"}`)
	err := json.Unmarshal(jsonData, &myJsonMap)
	if err != nil {
	panic(err)
	}
	fmt.Printf("%s\n", myJsonMap["hello"])
```

# objects
Composed of another objects or types.

# pointers 

# functions 
define functions as variables and pass them to other functions
## Anonymous function: assigned function to a variable
```go
func main(){
 add := func(m int){ return m+1 }
 result := add(6)
 println(result)
}
```
## Closures:
A closure is a function value that references variables from outside its body. 

The function may access and assign to the referenced variables; 
in this sense the function is "bound" to the variables.

For example, the adder function returns a closure. Each closure is bound to its own sum variable.
```go
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
```

# interface 