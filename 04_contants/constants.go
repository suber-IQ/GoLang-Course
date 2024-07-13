package main

import "fmt"

const age = 30

// name := "golang" // expected declaration, found nam
// var name string = "golang"

func main() {
	//:=
	// const name = "golang"

	// name = "javascript" //not changeable

	// const age = 30
	// fmt.Println(age)

	const (
		port = 5000
		host = "localhost"
	)

	// port = 8000 //cannot assign to port (neither addressable nor a map index expression)

	fmt.Println(port, host)

}
