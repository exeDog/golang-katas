package main

func main() {
	cards := NewDeck()

	hand, remainingCards := cards.Deal(5)

	hand.Print()
	remainingCards.Print()

	err := cards.SaveToFile("cardFile")
	if err != nil {
		return
	}

	createdFile, _ := NewDeckFromFile("cardFile")

	createdFile.Print()
}
