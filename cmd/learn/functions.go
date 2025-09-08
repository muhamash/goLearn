package main

import (
	"fmt"
)

func functions(){
	printLine("test")

	const var1 int = 11
	const var2 int = 1
	// var result float32 = divisionValue(var1, var2)
	// fmt.Println(result)


	var intDivisionValue, reminder , err  = intDivision(var1, var2)
	// if err != nil{
	// 	fmt.Println(err, err.Error())
	// }else{

	// 	fmt.Printf("Lets check that printf function with values %v and %v", intDivisionValue, reminder)
	// }

	switch{
		case err != nil:
			fmt.Println(err, err.Error())
		case reminder == 0:
			// fmt.Printf("Value %v and integer division: ", intDivisionValue)
			fmt.Println(intDivisionValue)
			fmt.Printf("Value %v and integer division\n", intDivisionValue)

		default:
			fmt.Printf("The result of %v and %v", intDivisionValue, reminder)		
	}

	switch reminder{
		case 0:
			fmt.Printf("The division was exact!")
		case 1,2:
			fmt.Println("The division was close!")
		default:
			fmt.Println("The division was not close!!")			
	}

	
}


func printLine(value string){
	fmt.Println("Hello world", value)
}

func divisionValue(valueOne int, valueTwo int) float32{
	var result float32 = float32(valueOne) / float32(valueTwo)
	return result
}

func intDivision(inputOne int, inputTwo int)(int, int, error){
	var err error
	// fmt.Println(err)
	if inputTwo == 0{
		// err = errors.New("can not have 0 value")
		err := fmt.Errorf("cannot divide %d by %d", inputOne, inputTwo)
		return  0, 0, err
	}
	// fmt.Println(err)

	var division int = inputOne/inputTwo
	var modValue int = inputOne%inputTwo
	
	return  division, modValue, err
}