package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d\n", sum(1, 2, 3))
	avg, err := avgFunc(4, 5, 6, 7, 8)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%.2f\n", avg)
	}
}

func sum(args ...int) (result int) { //Function with undetermined number of parameters
	for _, v := range args {
		result += v
	}
	return
}

func avgFunc(args ...int) (res float32, err error) {
	res = 0
	count, v := 0, 0
	for count, v = range args {
		res = res + float32(v)
	}
	res = res / float32(count)
	//todo
	return
}
