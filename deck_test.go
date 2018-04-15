package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first element of deck to be 'Ace of Spades', got %s", d[0])
	}

	if d[len(d) - 1] != "King of Clubs" {
		t.Errorf("Expected last element of deck to be 'King of Clubs', got %s", d[len(d) - 1])
	}

}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	err := d.saveToFile("_decktesting")
	if err != nil {
		t.Errorf("Couldn't save deck to file")
	}
	fileDeck := newDeckFromFile("_decktesting")
	if len(d) != len(fileDeck) {
		t.Errorf("Length mismatch between loaded deck and saved deck")
	}
	os.Remove("_decktesting")
}
