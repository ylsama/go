package main

func main() {
	ten := 10 //inferred
	if ten == 20 {
		println("This shouldn't be printed as 10 isn't equal to 20")
	} else if 11 == 11 && "go" == "go" {
		println("This isn't print because previous condition was satisfied")
	} else {
		println("In case no condition is satisfied, print this")
	}
}
