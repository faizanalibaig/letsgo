package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be 'Ace of Spades', but got %s", d[0])
	} else if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be 'Four of Clubs', but got %s", d[len(d)-1])
	}
}

func TestDeckSaveToFile (t *testing.T) {
	os.Remove("_save_deck")
	
	d := NewDeck()
	err := d.saveToFile("_save_deck")

	if err != nil {
		t.Errorf("Error saving deck to file: %v", err)
	}

	deck := readDeckFile("_save_deck")

	if len(deck) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(deck))
	}

	os.Remove("_save_deck")
}