package main

import (
	"fmt"
)

// slice -> dynamic array
// most used construct in go
// + useful methods
func main() {

	// var nums []int // uninitialized slice is nil (according other programing like null)
	// fmt.Println(nums)        // []
	// fmt.Println(nums == nil) // true
	// fmt.Println(len(nums)) // 0

	// 🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟👇
	// var nums = make([]int, 2) // [0 0] (initial size of array is 2)
	// var nums = make([]int, 2) // 2 ( capacity of array (if not given by third argu(capacity) then take length of array to asign with capacity))
	// var nums = make([]int, 2, 5) // 5 (capacity of array (third arg))

	// capacity -> maximum numbers of elements can fit
	// fmt.Println((cap(nums)))
	// fmt.Println(nums)
	// fmt.Println(nums == nil) // false

	// 🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟👇
	// var nums = make([]int, 2, 5)

	// nums = append(nums, 1) // add element ([0 0 1])
	// nums = append(nums, 2) // add element ([0 0 1 2]) // capacity: 5
	// nums = append(nums, 3) // add element ([0 0 1 2 3]) // capacity: 5
	// nums = append(nums, 4) // add element ([0 0 1 2 3 4]) // capacity: 10 (if not available capicity then double capacity)
	// nums = append(nums, 5) // add element ([0 0 1 2 3 4 5])
	// nums = append(nums, 6)  // add element ([0 0 1 2 3 4 5 6])
	// nums = append(nums, 7)  // add element ([0 0 1 2 3 4 5 6 7])
	// nums = append(nums, 8)  // add element ([0 0 1 2 3 4 5 6 7 8])
	// nums = append(nums, 9)  // add element ([0 0 1 2 3 4 5 6 7 8 9])
	// nums = append(nums, 10) // add element ([0 0 1 2 3 4 5 6 7 8 9 10]) // capacity: 20 (if not available capicity then double capacity)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))

	// 🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟👇
	// var nums = make([]int, 0, 5)
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)

	// fmt.Println(nums)
	// fmt.Println(cap(nums))

	// 🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟👇👆
	// nums := []int{}

	// fmt.Println(nums)  // []
	// fmt.Println(cap(nums)) // 0
	// fmt.Println(len(nums)) // 0

	// 🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟👇👆
	// nums := []int{}

	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// fmt.Println(nums)      // [1 2]
	// fmt.Println(cap(nums)) // 2
	// fmt.Println(len(nums)) // 2

	// 🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟👇👆

	// var nums = make([]int, 2, 5)
	// nums[0] = 5  // output:  [5 0]
	// fmt.Println(nums)

	// 🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟👇👆

	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2) // [2]

	// var nums2 = make([]int, len(nums))
	// // fmt.Println(nums, nums2) // [2] []

	// copy(nums2, nums) // output: [2] [2]
	// //copy function
	// fmt.Println(nums, nums2) // [2] []

	// 🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟👇👆
	// slice operator

	// var nums = []int{1, 2, 3}
	// fmt.Println(nums[0:1]) // [1]
	// fmt.Println(nums[0:2]) // [1 2]
	// fmt.Println(nums[:2])  // [1 2]
	// fmt.Println(nums[1:])  // [2 3]

	// 🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟👇👆
	// slices package
	// var nums1 = []int{1, 2}
	// var nums2 = []int{1, 2}
	// fmt.Println(slices.Equal(nums1, nums2)) //output true

	// var nums1 = []int{1, 2, 3}
	// var nums2 = []int{1, 2, 4}

	// fmt.Println(slices.Equal(nums1, nums2)) //output false

	// 🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟👇👆

	var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums) //[[1 2 3] [4 5 6]]

}
