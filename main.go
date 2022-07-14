package main

func main() {
	cards := newDeck()
	pointerToDeck := &cards
	pointerToDeck.addCard(card{Value: "R", Suit: "♡"})
	cards.saveToFile("my_cards")
	readFromFile("my_cards")
	cards.shuffle()
	cards.print()
}
