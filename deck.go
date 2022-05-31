package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of Deck
// wich is a slice of strings
type deck []string

func newDeck() deck {
	suits := []string{"♡", "♠", "♢", "♣"}
	values := deck{"2", "3", "4", "5", "6", "7"}
	cards := deck{}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+suit)
		}
	}

	return cards
}

func (d deck) print() {

	//Cycle
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func readFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		// os.Exit(1)
		return newDeck()
	}

	s := strings.Split(string(bs), ",")
	d := deck(s)

	fmt.Println(s)

	return d
}
