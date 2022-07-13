package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 28 {
		t.Errorf("Expected deck length of 28, but got %d", len(d))
	}

	if d[0] != "2♡" {
		t.Errorf("Expected first card of 2♡, but got %s", d[0])
	}

	if d[len(d)-1] != "8♣" {
		t.Errorf("Expected last card of 2♡, but got %s", d[len(d)-1])
	}
}

func TestNewDeckSaveToFileAndReadFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := readFromFile("_decktesting")

	if len(loadedDeck) != 28 {
		t.Errorf("Expected deck length of 28, but got %d", len(loadedDeck))
	}

	if loadedDeck[0] != "2♡" {
		t.Errorf("Expected first card of 2♡, but got %s", loadedDeck[0])
	}

	if loadedDeck[len(loadedDeck)-1] != "8♣" {
		t.Errorf("Expected last card of 2♡, but got %s", d[len(loadedDeck)-1])
	}

	os.Remove("_decktesting")
}
