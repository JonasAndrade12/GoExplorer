package main

func main() {
	cards := newDeck()
	cards.addCard(card{Value: "R", Suit: "♡"})
	cards.saveToFile("my_cards")
	readFromFile("my_cards")
	cards.shuffle()
	cards.print()
}
