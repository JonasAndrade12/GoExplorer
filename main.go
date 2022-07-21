package main

import "fmt"

func main() {
	cards := newDeck()
	cards.addCard(card{Value: "R", Suit: "♡"})
	cards.saveToFile("my_cards")
	readFromFile("my_cards")
	cards.shuffle()
	cards.print()

	// Learn Map
	colors := make(map[string]string)
	colors["white"] = "#fff"
	colors["black"] = "#000"
	delete(colors, "white")
	fmt.Println(colors)
}
