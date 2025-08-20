package utils

import "fmt"

func SliceInteger() {
	var slice = []int{50, 75, 66, 20, 32, 90}
	var index uint

	for k, s := range slice {
		if s == 66 {
			index = uint(k)
		}
	}

	slice = append(slice[:index+1], append([]int{88}, slice[index+1:]...)...)

	fmt.Println(slice)
}
