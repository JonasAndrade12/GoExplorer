package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d.Cards) != 28 {
		t.Errorf("Expected deck length of 28, but got %d", len(d.Cards))
	}

	if d.Cards[0].Suit != "♡" && d.Cards[0].Value != "2" {
		t.Errorf("Expected first card of 2♡, but got %s", d.Cards[0])
	}

	if d.Cards[len(d.Cards)-1].Suit != "♣" && d.Cards[len(d.Cards)-1].Value != "8" {
		t.Errorf("Expected last card of 8♣, but got %s", d.Cards[len(d.Cards)-1])
	}
}

func TestNewDeckSaveToFileAndReadFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := readFromFile("_decktesting")

	if len(loadedDeck.Cards) != 28 {
		t.Errorf("Expected deck length of 28, but got %d", len(loadedDeck.Cards))
	}

	if loadedDeck.Cards[0].Suit != "♡" && loadedDeck.Cards[0].Value != "2" {
		t.Errorf("Expected first card of 2♡, but got %s", loadedDeck.Cards[0])
	}

	if loadedDeck.Cards[len(loadedDeck.Cards)-1].Suit != "♣" && loadedDeck.Cards[len(loadedDeck.Cards)-1].Value != "8" {
		t.Errorf("Expected last card of 8♣, but got %s", loadedDeck.Cards[len(loadedDeck.Cards)-1])
	}

	os.Remove("_decktesting")
}

func TestAddCardToDeck(t *testing.T) {
	d := newDeck()

	if len(d.Cards) != 28 {
		t.Errorf("Expected deck length of 28, but got %d", len(d.Cards))
	}

	d.addCard(card{Value: "R", Suit: "♡"})

	if len(d.Cards) != 29 {
		t.Errorf("Expected deck length of 29, but got %d", len(d.Cards))
	}

	if d.Cards[len(d.Cards)-1].Suit != "♡" && d.Cards[len(d.Cards)-1].Value != "R" {
		t.Errorf("Expected first card of R♡, but got %s", d.Cards[0])
	}
}
