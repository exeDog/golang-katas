package cards

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// Deck Create a new type of 'deck' which is a slice of strings
type Deck []string

func NewDeck() Deck {
	cards := Deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}

	return cards
}

func (d Deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d Deck) Deal(handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

func (d Deck) toString() string {
	return strings.Join(d, ",")
}

func (d Deck) SaveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func NewDeckFromFile(fileName string) (Deck, error) {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		NewDeck()
	}

	return toDeck(bs), nil

}

func toDeck(bytes []byte) Deck {
	return Deck(strings.Split(string(bytes), ","))
}

func (d Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
}
