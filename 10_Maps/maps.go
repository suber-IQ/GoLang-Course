package main

import (
	"fmt"
	"maps"
)

// maps -> hash,object,dictionary

func main() {
	// 👉 🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟
	//creating map

	// m := make(map[string]string)

	// setting an element
	// m["name"] = "golang"
	// m["area"] = "backend"

	// get a element
	// fmt.Println((m["name"])) // golang
	// fmt.Println((m["area"])) // backend
	// if don't exists key on map  (IMP: if key does not exist in the map then it returns zero value)
	// fmt.Println((m["phone"])) //    (empty value)

	// 👉 🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟
	// 🎄 IMP: all data type with different zero values int:0 , string:empty string, bool:false like to other...
	// m := make(map[string]int)
	// m["age"] = 30
	// fmt.Println(m["age"])   // 30
	// fmt.Println(m["phone"]) // 0 (IMP: if key does not exist in the map then it returns zero value)

	// 👉 🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟
	//   find length
	// m := make(map[string]int)
	// m["age"] = 30
	// m["price"] = 30
	// fmt.Println(len(m)) // 2

	// 👉 🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟
	//   delete and clear
	// m := make(map[string]int)
	// m["age"] = 30
	// m["price"] = 30

	// delete(m, "price")  // output delete after print len: 1 , delete element m to price key field value
	// fmt.Println(len(m)) //output delete after print len: 1
	// clear(m)            // all element is clear
	// fmt.Println(len(m)) // output after clear len: 0

	// // 👉 🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟
	// m := map[string]int{"price": 30, "age": 20}
	// fmt.Println(m) // output: map[age:20 price:30]

	// 👉 🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟
	// element get
	// m := map[string]int{"price": 30, "age": 20}
	// _, ok := m["price"] // output: all ok
	// _, ok := m["age"] // output: all ok
	// v, ok := m["phone"] // output: not ok
	// v, ok := m["phone"] // output: 0 (zero)
	// v, ok := m["age"] // output: 20
	// fmt.Println(v)
	// if ok {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("not ok")
	// }

	// 👉 🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟🌟
	//check equal
	// m1 := map[string]int{"price": 30, "age": 20}
	// m2 := map[string]int{"price": 30, "age": 20}

	// fmt.Println(maps.Equal(m1, m2)) //output: true

	m1 := map[string]int{"price": 30, "age": 20}
	m2 := map[string]int{"price": 30, "age": 35}

	fmt.Println(maps.Equal(m1, m2)) //output: false

}
