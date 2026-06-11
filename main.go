package main

import "fmt"

func anagramCheck(word1 string, word2 string) bool {
	
	return true
}

func main() {
	var word1 = "listen"
	var word2 = "silent"
	isAnagram := anagramCheck(word1, word2)

	if isAnagram == true {
		fmt.Println("It's Anagram")
	} else {
		fmt.Println("It's not Anagram")
	}
}
