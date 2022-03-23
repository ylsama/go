# Function
Functions can call other (or self) functions

## Boilderplate function
```go
func [function_name] (param1 type, param2 type...) (returned type1, returned type2...) {
 //Function body
}
```
### 1. example 1 (no return type)
```go
func hello(message string) {
 fmt.Printf("Hello %s\n", message)
 return nil
}
```
### 2. example 2 (with return type)
```go
func hello(x int8) int8 { return x + 7 }
```
### 3. example 3 (with return variable) Naming returned types
```go
func doesReturnError(arg1 int) (x int, err error) {
	x = arg1 + 10 //already declare
	err = errors.New("this function simply returns an error")
	return //no need to specify the return x, err
}
```
## Function with undetermined number of parameters
```go
func sum(args ...int) (result int) {
    for _, v := range args {
        result += v
    }
    return
}
```
# Function as variable value: 
## Anonymous function: 
assigned function to a variable
```go
func main(){
 add := func(m int){ return m+1 }
 result := add(6)
 //1 + 6 must print 7
 println(result)
}
```

# closures
```go
func main(){
    addn := func(m int){
        return func(n int){ return m+n }
    }
    addfive := addn(5)
    addsix := addn(6)
    println(addfive(2))
    println(addsix(1))
}
```
