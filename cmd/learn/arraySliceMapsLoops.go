package main

import "fmt"

func arraySliceMapsLoops(){
	fmt.Printf("Learning array maps slice loops\n \n")

	// arr := [5]int{10, 20, 30, 40, 50}
    // for i := 0; i < len(arr); i++ {
    //     fmt.Printf("arr[%d] = %d, address = %p\n", i, arr[i], &arr[i])
    // }

	// var intArr [3]int32 = [3]int32{1,3,4}
	// intArr := [4]int{1,3,56}
	
	// intArr[3] = 23
	// // intArr.append()
	// fmt.Println(intArr, intArr[1:2], &intArr[1], intArr[2:2])

	// var sliceArr []int = []int{1,2,3,4,5,7}
	// fmt.Println(sliceArr, cap(sliceArr))
	// sliceArr = append(sliceArr, 23,2345)
	// fmt.Println(sliceArr)
	// sliceArr = append(sliceArr, intArr[2])
	// fmt.Println(sliceArr, cap(sliceArr))

	// var newArr []int = []int{1,3}
	// sliceArr = append(sliceArr, newArr...)
	// fmt.Println(sliceArr, cap(sliceArr))

	// var makSliceArr []int = make([]int, 12, 30)
	// fmt.Println(makSliceArr)

	//--->> maps

	// var firstMap map[string]uint = make(map[string]uint)
	var myMap = map[string]uint{"ash": 26, "awish": 20}
	// fmt.Println(firstMap, myMap, myMap["ash"], firstMap["test"])

	// var age, ok = myMap["awish"]
	// if !ok{
	// 	fmt.Printf("no value")
	// }else{
	// 	fmt.Printf("age: %v", age)
	// }


	var newMap = map[string]uint{"ash": 26, "awish": 22, "dad": 55, "mom": 45}
	for name := range myMap{
		fmt.Printf("Name: %v and value :%v\n", name, myMap[name])
	}

	for name, age := range newMap{
		fmt.Println(name, age)
	}

	var i int = 0
	for i < 10{
		fmt.Println(i)
		i++
	}

	j := 0
	for{
		if j>=5{
			break
		}
		fmt.Println(i, j)
		j++
	}

}