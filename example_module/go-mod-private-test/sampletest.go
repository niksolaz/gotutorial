package sampletest

import "fmt"

func ReadTest(t string) (result string) {
	result = "WOW " + t
	fmt.Println(result)
	return result
}
