package main

type deck []string

func newDeck() deck {
	card := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			card = append(card, value+" of "+suit)
		}
	}

	return card
}
