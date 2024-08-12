package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type deck []card
type playerHand []card
type houseHand []card

func newDeck() deck {
	decks := deck{}
	cardSuits := []string{"Spade", "Clubs", "Diamond", "Hearts"}
	cardValues := []string{"Ace", "Jack", "Queen", "King", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cardString := value + " of " + suit
			cardValue, _ := convertCardValue(value)

			decks = append(decks, createCard(cardString, cardValue))
		}
	}
	decks.shuffleDeck()
	return decks
}

func (d deck) shuffleDeck() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i, _ := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func convertCardValue(value string) (int, error) {
	switch value {
	case "Ace":
		return 11, nil
	case "Jack", "Queen", "King":
		return 10, nil
	default:
		v, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("something went wrong")
		}

		return v, err
	}
}

func (d deck) dealCards(p deck, h deck) (deck, deck, deck, string, string) {
	houseText := "Dealer Hand: "
	playerText := "Player Hand: "

	for i := 0; i < 4; i++ {
		if i%2 == 0 {
			if i == 2 {
				houseText += "*"
				// cardDetail := d[i].getCard()
				// houseText += cardDetail + ", "
			} else {
				cardDetail := d[i].getCard()
				houseText += cardDetail + ", "
			}
			h = append(h, d[i])
		} else {
			p = append(p, d[i])
			cardDetail := d[i].getCard()
			playerText += cardDetail + ", "
		}
	}
	d = d[4:]

	return d, p, h, playerText, houseText
}

func (d deck) hitCards(x deck) (deck, deck) {
	cardToAdd := d[0]
	x = append(x, cardToAdd)
	d = d[1:]

	return d, x
}

func getValue(d deck) int {
	deckValue := 0
	aceCount := []card{}
	normalCount := []card{}
	for _, card := range d {
		cType := strings.Split(card.cardType, " ")[0]
		if cType == "Ace" {
			aceCount = append(aceCount, card)
		} else {
			normalCount = append(normalCount, card)
		}
	}
	for _, card := range normalCount {
		deckValue += card.cardValue
	}
	if len(aceCount) > 1 {
		deckValue += 1 * len(aceCount)
	} else if len(aceCount) == 1 {
		if deckValue+11 > 21 {
			deckValue += 1
		} else {
			deckValue += 11
		}
	}
	return deckValue

}
