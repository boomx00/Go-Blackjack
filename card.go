package main

type card struct {
	cardType  string
	cardValue int
}

func createCard(ct string, v int) card {
	return card{cardType: ct, cardValue: v}
}

func (c card) getValue() int {
	return c.cardValue
}

func (c card) getCard() string {
	return c.cardType
}
