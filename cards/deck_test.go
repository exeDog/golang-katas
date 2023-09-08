package cards

import (
	"os"
	"testing"
)

func TestNewDeckTest(t *testing.T) {
	d := NewDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but instead got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Four of Clubs, but instead got %v", d[len(d)-1])
	}
}

func TestDeck_SaveToFileAndNewDeckFromFile(t *testing.T) {
	const FILENAME = "_decktesting"

	err := os.Remove(FILENAME)
	if err != nil {
		return
	}

	deck := NewDeck()

	err = deck.SaveToFile(FILENAME)
	if err != nil {
		return
	}

	loadedDeck, err := NewDeckFromFile(FILENAME)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadedDeck))
	}
}
