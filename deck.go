package main

import "fmt"

// Create a new type of Deck
// wich is a slice of strings
type deck []string

func (d deck) print() {

	//Cycle
	for i, card := range d {
		fmt.Println(i, card)
	}
}
