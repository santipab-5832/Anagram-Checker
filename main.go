package main

import "fmt"

func anagramCheck(word1 string, word2 string) bool {
	count := 0
	if len(word1) == len(word2) {
		runes1 := []rune(word1)
		runes2 := []rune(word2)
		for _, v1 := range runes1 {
			for _, v2 := range runes2 {
				if v1 == v2 {
					count++
				}
			}
		}
		checkWord := len(word1)
		if count == checkWord {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
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
