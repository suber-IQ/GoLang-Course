package main

//numbered sequence of specific length (same type)

func main() {

	// 👉 zeroed values
	// int -> 0, string -> "", bool -> false
	// var nums [4]int

	// nums[0] = 1

	// fmt.Println(nums[0]) // 1 (get element)
	// fmt.Println(nums)    // [1 0 0 0]
	//👉 array length
	// fmt.Println(len(nums)) // 4

	// 👉 boolean arrays
	// var vals [4]bool // [false false false false]
	// vals[2] = true   //[false false true false]
	// fmt.Println(vals)

	// 👉 string arrays
	// var name [3]string // empty string
	// name[0] = "golang"

	// fmt.Println(name)

	// 👉 declare with value in single line
	// nums := [3]int{1,2,3,4}  // 4 with error:  (index 3 is out of bounds (>= 3))
	// nums := [3]int{1, 2, 3}

	// fmt.Println(nums)

	// 👉 2d arrays
	// nums := [2][2]int{{3, 4}, {5, 6}} // output: [[3 4] [5 6]]

	// fmt.Println(nums)

	//  👉 Point here 👇
	// - fixed size, that is predictable
	// - Memory optimazation
	// - Contant time access

}
