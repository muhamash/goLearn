package main

import (
	"fmt"
	"goLearn/cmd/code"
)

func main(){
	
	wordlist1 := []string{"KiTe", "kite", "hare", "Hare"}
	queries1 := []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"}
	fmt.Println(code.SpellChecker(wordlist1, queries1))
	// Output: ["kite","KiTe","KiTe","Hare","hare","","","KiTe","","KiTe"]

	// Example 2
	wordlist2 := []string{"yellow"}
	queries2 := []string{"YellOw"}
	fmt.Println(code.SpellChecker(wordlist2, queries2))
	// Output: ["yellow"]


}
