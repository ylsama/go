package main

func main() {
	ifElseFlow()
	stwichFlow()
	forFlow()
}

func ifElseFlow() {
	ten := 10 //inferred
	if ten == 20 {
		println("This shouldn't be printed as 10 isn't equal to 20")
	} else if 11 == 11 && "go" == "go" {
		println("This isn't print because previous condition was satisfied")
	} else {
		println("In case no condition is satisfied, print this")
	}
}

func stwichFlow() {
	var (
		number int8 = 3
	)
	switch number {
	case 1:
		println("Number is 1")
	case 2:
		println("Number is 2")
	case 3:
		println("Number is 3")
	}
}

func forFlow() {
	var (
		tmp int8 = 0
	)
	for i := 0; i < 10; i++ {
		tmp++
	}
	println(tmp)
}
