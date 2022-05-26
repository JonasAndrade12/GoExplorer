package main

import "fmt"

func main() {
	card := "2 ♥"
	card2 := newCard()
	fmt.Println(card)
	fmt.Println(card2)
}

func newCard() string {
	return "5 ♢"
}
