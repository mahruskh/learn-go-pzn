package main

import (
	"fmt"
	"strings"
)

func sayHello(name string) string {
	return "Hello " + name
}

// To short way, using type
type Filter func(string) string

// func sayHelloWithFilter(name string, filter func(string) string) string {
func sayHelloWithFilter(name string, filter Filter) string {
	return "Hello, " + filter(name)
}

func isBlockedWord(value string) bool {
	words := []string{
		"anjing", "goblok", "monyet", "asu",
	}
	value = strings.ToLower(value)
	for _, word := range words {
		if strings.Contains(value, word) {
			return true
		}
	}
	return false
}

func spamFilter(sentence string) string {
	words := strings.Split(sentence, " ")
	for index, word := range words {
		if isBlockedWord(word) {
			words[index] = strings.Repeat("*", len(word))
		}
		// fmt.Println(words)
	}
	return strings.Join(words, " ")
}

func spaceCode() {
	fmt.Println("-------------------------------------------------------------")
}

func main() {
	// Funcation value
	// Function as data typ, can assign to a variable
	fmt.Println(sayHello("Mahrus basic func"))
	sayHandler := sayHello
	fmt.Println(sayHandler("Khomaini from func value"))

	spaceCode()

	// Function as parameter
	sentence := "Anda anjing, Memang Monyet kamu asu lah."
	fmt.Println(sentence)

	result := sayHelloWithFilter(sentence, spamFilter)
	fmt.Println(result)

}
