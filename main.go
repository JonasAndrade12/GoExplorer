package main

func main() {
	cards := deck{newCard(), newCard(), "2 ♥"}

	//Add element
	cards = append(cards, "4 ♣")

	// Show cards
	cards.print()
}

func newCard() string {
	return "5 ♢"
}
