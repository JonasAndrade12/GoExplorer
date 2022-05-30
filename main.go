package main

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	// Show cards
	hand.print()
	remainingDeck.print()
}
