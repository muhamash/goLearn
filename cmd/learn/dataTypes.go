package main

import (
	"fmt"
	"unicode/utf8"
)

func dataTypes (){
	fmt.Println("Hello world")

	var intNum uint64 = 4234234234232342334
	intNum = intNum + 3456345234534345323
	fmt.Println(intNum)

	var floatNum float32 = 23423.2323452345
	fmt.Println(floatNum)

	var floatNum32 float32 = 53.1021012
	var intNum32 int64 = 2323784562378
	var resultOne = floatNum32 + float32(intNum32)
	// var resultTwo = floatNum32 + float32(intNum32)
	fmt.Println(resultOne)

	var inputOne int = 5
	var inputTwo int = 2
	fmt.Println(inputOne/inputTwo)
	fmt.Println(inputOne%inputTwo)
	
	const stringValue string = "hey" + "  " + "test"
	fmt.Println(stringValue, len(stringValue), len("Y"), len("y"), len("A"), len("Q"), len("a"), len("q"), len("𝒴"), utf8.RuneCountInString("𝒴"))

	var stringRune rune = 'a'
	fmt.Println(stringRune)
	fmt.Println(string(stringRune))
	fmt.Println(len("ক"))     
	fmt.Println(utf8.RuneCountInString("ক"))

	// const myVar string = foo();
	// fmt.Println(myVar)
}