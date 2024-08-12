package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type balance struct {
	amount int
}

func (bPointer *balance) updateBalanceBet(v int) {
	(*bPointer).amount -= v
}

// func (b balance) winBalance(change int) int {
// 	finalBalance := b.amount + change
// 	return finalBalance
// }

func (bPointer *balance) updateBalanceWinOrDraw(v int) {
	(*bPointer).amount += v
}

func main() {
	balance := balance{amount: 100}
	reader := bufio.NewReader(os.Stdin) // Create a new reader to read from standard input (terminal)

	fmt.Println("Would you like to play a game of blackjack: ")
	fmt.Print("1. Yes \n2. No") // Prompt the user for input
	fmt.Println()
	choice, _ := reader.ReadString('\n')
	playGame := false
	choice = strings.TrimSpace(choice)
	if choice == "1" {
		playGame = true
	} else {
		os.Exit(1)
	}

	for playGame {
		decks := newDeck()
		betAmount := bet()
		if betAmount > 0 {
			var playerHand, houseHand deck
			var playerString, houseString string
			balance.updateBalanceBet(betAmount)
			fmt.Println("Current balance ", balance.getBalance())
			time.Sleep(time.Second)
			fmt.Println("Cards Being Delt")
			decks, playerHand, houseHand, playerString, houseString = decks.dealCards(playerHand, houseHand)
			if getValue(houseHand) == 21 {
				fmt.Println("You Lost, Starting new game")
			} else {
				printHands(houseString)
				printHands(playerString)
				shChoice := standOrHit()
				// 1 for stand
				if shChoice == "1" {
					houseValue := revealCards(houseHand, "house")
					playerValue := revealCards(playerHand, "player")

					for houseValue < 17 {
						decks, houseHand = decks.hitCards(houseHand)
						houseValue = getValue(houseHand)
					}
					fmt.Print(houseHand)
					fmt.Println(houseValue)
					fmt.Print(playerHand)
					fmt.Println(playerValue)
					if houseValue > 21 {
						fmt.Println("You Win!")
						balance.updateBalanceWinOrDraw(betAmount * 2)
						fmt.Println("New Balance: ", balance.amount)
						fmt.Println("Starting new game")
					} else if playerValue > houseValue {
						fmt.Println("You Win!")
						balance.updateBalanceWinOrDraw(betAmount * 2)
						fmt.Println("New Balance: ", balance.amount)
						fmt.Println("Starting new game")
					} else if playerValue == houseValue {
						fmt.Println("Draw")

						balance.updateBalanceWinOrDraw(betAmount)
					} else {
						fmt.Println("You Lose, New Balance: ", balance.amount)
						fmt.Println("Starting new game")

						// tryAgain()
					}
				}
				// 2 for hit
				if shChoice == "2" {
					decks, playerHand = decks.hitCards(playerHand)
					playerValue := revealCards(playerHand, "player")
					fmt.Println(houseHand)
					fmt.Println(playerHand)
					fmt.Println(playerValue)
					if playerValue > 21 {
						fmt.Println("You Lose, New Balance: ", balance.amount)
						fmt.Println("Starting new game")
					} else {
						for playerValue <= 21 {
							hitChoice := hitPrompt()
							if hitChoice == "1" {
								decks, playerHand = decks.hitCards(playerHand)
								fmt.Print(houseHand)
								fmt.Println(playerHand)
								playerValue = getValue(playerHand)
							} else if hitChoice == "2" {
								houseValue := revealCards(houseHand, "house")
								playerValue := revealCards(playerHand, "player")

								for houseValue < 17 {
									decks, houseHand = decks.hitCards(houseHand)
									houseValue = getValue(houseHand)
								}
								fmt.Print(houseHand)
								fmt.Println(houseValue)
								fmt.Print(playerHand)
								fmt.Println(playerValue)
								if houseValue > 21 {
									fmt.Println("You Win!")
									balance.updateBalanceWinOrDraw(betAmount * 2)
									fmt.Println("New Balance: ", balance.amount)
									fmt.Println("Starting new game")
								} else if playerValue > houseValue {
									fmt.Println("You Win!")
									balance.updateBalanceWinOrDraw(betAmount * 2)
									fmt.Println("New Balance: ", balance.amount)
									fmt.Println("Starting new game")
								} else if playerValue == houseValue {
									fmt.Println("Draw")
									balance.updateBalanceWinOrDraw(betAmount)
								} else {
									fmt.Println("You Lose, New Balance: ", balance.amount)
									fmt.Println("Starting new game")

									// tryAgain()
								}
								break
							}
						}
					}

				}
			}

		}
	}

}

func (b balance) getBalance() int {
	return b.amount
}

func bet() int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("How much would you like to bet")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		} else {
			return number
		}
	}
}

func standOrHit() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Stand or Hit?")
		fmt.Println("1. Stand")
		fmt.Println("2. Hit")
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
		} else {
			return input
		}
	}
}

func revealCards(d deck, s string) int {
	returnText := ""
	if s == "player" {
		returnText = "Player Deck: "
	} else {
		returnText = "Dealer Deck: "
	}

	returnValue := 0
	for _, card := range d {
		returnText += card.cardType + ", "
		returnValue += card.cardValue
	}
	return returnValue
}

func printHands(s string) {
	fmt.Println(s)
}

func tryAgain() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Try again?")
		fmt.Println("1. yes")
		fmt.Println("2. no")
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
		} else {
			return input
		}
	}
}

func hitPrompt() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Hit again?")
		fmt.Println("1. yes")
		fmt.Println("2. no")
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
		} else {
			return input
		}
	}
}
