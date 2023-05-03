package main

import "fmt"

func main() {
	array1 := [3]string{"a", "c", "d"}
	array2 := [3]string{"a", "d", "c"}

	for i := 0; i < len(array1); i++ {
		if array1[i] != array2[i] {
			fmt.Printf("Index ke %d berbeda\n", i+1)
		}
	}
}
