package main

import "fmt"

//for -> only construct in go for looping

func main() {

	//while loop
	// i := 1

	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	//infinite loop

	// for {
	// 	println("1")
	// }

	//classic for loop
	// for i := 0; i < 3; i++ {
	// 	fmt.Println(i)
	// }

	//break and continue for loop
	// for i := 0; i < 3; i++ {
	// 	// fmt.Println(i)

	// 	// break
	// 	if i == 1 { // 1 is skip
	// 		continue
	// 	}
	// 	fmt.Println(i)

	// }

	// range

	for i := range 11 {
		fmt.Println(i)
	}

}
