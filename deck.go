package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

// struct of a card
type card struct {
	Value string
	Suit  string
}

// Create a new type of Deck
// wich is a slice of strings
type deck struct {
	Cards []card
	Title string
}

func newDeck() deck {
	suits := []string{"♡", "♠", "♢", "♣"}
	values := []string{"2", "3", "4", "5", "6", "7", "8"}
	cards := deck{}

	for _, suit := range suits {
		for _, value := range values {
			cards.Cards = append(cards.Cards, card{Value: value, Suit: suit})
		}
	}

	return cards
}

func (d deck) print() {
	//Cycle
	for i, card := range d.Cards {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {

	out, err := json.Marshal(d)

	if err != nil {
		panic(err)
	}

	return string(out)
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func readFromFile(filename string) deck {
	var d deck

	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		// os.Exit(1)
		return newDeck()
	}

	e := json.Unmarshal(bs, &d)

	if e != nil {
		panic(e)
	}

	return d
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d.Cards {
		newPosition := r.Intn(len(d.Cards) - 1)

		d.Cards[i], d.Cards[newPosition] = d.Cards[newPosition], d.Cards[i]
	}
}

func (pointerToDeck *deck) addCard(card card) {
	(*pointerToDeck).Cards = append((*pointerToDeck).Cards, card)
}
