package main

import "fmt"

func main() {
	cards := []string{newCard(), newCard(), "2 ♥"}

	//Add element
	cards = append(cards, "4 ♣")

	//Cycle
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "5 ♢"
}
