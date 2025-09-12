package main

import (
	"fmt"
	"goLearn/cmd/code"
)

func main(){
	// nums := []int{1,3,5,6}

    fmt.Println(code.SearchInsert([]int{1, 3, 5, 6}, 5)) // Output: 2
	fmt.Println(code.SearchInsert([]int{1, 3, 5, 6}, 2)) // Output: 1
	fmt.Println(code.SearchInsert([]int{1, 3, 5, 6}, 7)) // Output: 4
	fmt.Println(code.SearchInsert([]int{1, 3, 5, 6}, 0)) // Output: 0)             


}
