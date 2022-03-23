package main

import "errors"

func main() {
	result, err := doesReturnError(1)
	if err != nil {
		panic(err)
	} else {
		print(result)
	}
}
func doesReturnError(arg1 int) (x int, err error) {
	x = arg1 + 10
	err = errors.New("this function simply returns an error")
	return
}
