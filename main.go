package main

import (
	"fmt"
	"goLearn/cmd/code"
)

func main(){
	nums := []int{3, 2, 2, 3}
    val := 3
    k := code.RemoveElement(nums, val)

    fmt.Println("k =", k)             
    fmt.Println("nums =", nums[:k])  

    nums2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
    val2 := 2
    k2 := code.RemoveElement(nums2, val2)

    fmt.Println("k2 =", k2)           
    fmt.Println("nums2 =", nums2[:k2])
}
