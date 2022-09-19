package main

import "reflect"

func Anagrams(word string, words []string) []string {
	// your code
	var anagrams []string
	mainWordMap := getWordMap(word)
	for _, testWord := range words {
		testWordMap := getWordMap(testWord)
		if reflect.DeepEqual(testWordMap, mainWordMap) {
			anagrams = append(anagrams, testWord)
		}
	}
	return anagrams
}

func getWordMap(word string) map[rune]int {
	wordMap := make(map[rune]int)
	for _, c := range word {
		if _, ok := wordMap[c]; ok {
			wordMap[c]++
		} else {
			wordMap[c] = 0
		}
	}
	return wordMap
}
