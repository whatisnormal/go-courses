package main

import "fmt"

func main() {
	//var card string = "Ace of Spaces"

	// Line above and below are the same semantically.
	// ':=' is only used when initializing a variable, not re-assigning.
	//card := "Ace of Spaces"

	// -> re-assign
	//card = "Five of Diamonds"

	card := newCard()

	fmt.Println(card)

	//Slice -> an array that can grow and shrink.
	//# cards := []string{newCard(), newCard(), "Ace of Diamonds"}
	//# cards := deck{newCard(), newCard(), "Ace of Diamonds"}

	// Add a string to the very end of the slide 'cards'
	//# cards = append(cards, "Six of Spades")

	// Instantiates a newdeck from deck.go

	cards := newDeck()
	// **** Read from file
	//filename := "my_cards"
	//cards.saveToFile(filename)

	// **** Read from file
	//cardsFromFile := newDeckFromFile(filename)
	//cardsFromFile.print()

	// **** Test file that does not exist
	//newDeckFromFile("fileNotExists")

	fmt.Println(cards.toString())

	//Returns 2 decks.
	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()

	// Iterate over all cards. Assign index and card for the range result.
	/*#
	for i, card := range cards {
		// Prints index and card. Both 'i' and 'card' must be used on the code.
		fmt.Println(i, card)
	}
	*/
	cards.print()

	fmt.Println("**********")

	// Cast a string to a byte slice
	greeting := "Hi there!"
	fmt.Println([]byte(greeting))

	// ****** Test shuffle ******
	shuffledCards := newDeck()
	shuffledCards.shuffle()
	shuffledCards.print()

}

func newCard() string {
	return "Five of Diamonds."
}
