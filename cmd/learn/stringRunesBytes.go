package main

import (
	"fmt"
	"strings"
)

func StringRunesBytes(){
	fmt.Printf("Hello string, runes, bytes\n")

	var stringText =  []rune("råsµnπç")
	var indexed = stringText[1]
	fmt.Println(indexed, stringText)
	fmt.Printf("%v value and %T type \n", indexed, indexed)

	for i, v := range stringText {
		fmt.Println(i,v)
	}
	fmt.Printf("\nThe len of stringText : %v \n", len(stringText))

	var myRune = 'a'
	fmt.Printf("\n rune is %v\n", myRune)

	var strSlice = []string{"s", "t", "r", "i", "n", "g"}
	var castString = ""

	// built in package in go
	var stringBuilder  strings.Builder
	// fmt.Println(stringBuilder)

	for i := range strSlice {
		castString = castString + strSlice[i]
		fmt.Println(i, castString, strSlice)

		stringBuilder.WriteString(strSlice[i])
		// fmt.Println(stringBuilder)
	}
	var catString = stringBuilder.String()

	fmt.Println(castString, catString)
}