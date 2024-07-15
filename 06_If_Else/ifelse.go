package main

import "fmt"

func main() {
	//
	// age := 18
	// age := 16

	// if age >= 18 {
	// 	fmt.Println("Person is an adult...")
	// } else {
	// 	fmt.Println("Person is not an adult...")
	// }

	// 👉

	// age := 10

	// if age >= 18 {
	// 	fmt.Println("Person is an adult...")
	// } else if age >= 12 {
	// 	fmt.Println("Person is teenager...")
	// } else {
	// 	fmt.Println("Person is a kid...")
	// }

	// 👉

	// role := "admin"
	// var hasPermission = true
	// var hasPermission = false

	// if role == "admin" || hasPermission {
	// 	fmt.Println("yes")
	// }

	// if role == "admin" && hasPermission {
	// 	fmt.Println("yes")
	// }

	//👉  we can declare a variable inside if construct

	if age := 5; age >= 18 {
		fmt.Println("Person is an adult", age)
	} else if age <= 8 {
		fmt.Println("Person is an kid", age)
	} else {
		fmt.Println("Person is an teenager", age)
	}

	//👉  go does not have ternary, you will have to use normal if else

}
