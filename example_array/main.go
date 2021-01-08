package main

import (
	"fmt"
)

func setArray(v interface{}, i int, arr *[4]interface{}) {
	arr[i] = v
}
func main() {
	var myarr [4]interface{}
	myarr[0] = "we"
	myarr[1] = "4re"
	myarr[2] = 234
	myarr[3] = true
	fmt.Println("My Array: ", myarr)
	setArray(44, 1, &myarr)
	fmt.Println("My Array: ", myarr)
	var myslice = myarr[1:3]
	fmt.Println("My Slice: ", myslice)
	fmt.Println("My Slice capacity: ", cap(myslice))
	fmt.Println("My Slice lenght: ", len(myslice))
	var emptyslice = make([]int, 5)
	fmt.Println(emptyslice)
	emptyslice[3] = 34
	fmt.Println(emptyslice)
	for i, v := range emptyslice {
		fmt.Println("index: ", i, " - value:", v)
	}
}
