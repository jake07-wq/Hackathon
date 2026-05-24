package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	// The deck has 12 cards divided evenly among 4 players (3 cards each)
	cardsPerPlayer := len(deck) / 4

	for i := 0; i < 4; i++ {
		// Calculate the start and end index for slicing the deck
		start := i * cardsPerPlayer
		end := start + cardsPerPlayer
		playerDeck := deck[start:end]

		// Print the player number and their respective cards
		fmt.Printf("Player %d: %d, %d, %d\n", i+1, playerDeck[0], playerDeck[1], playerDeck[2])
	}
}
