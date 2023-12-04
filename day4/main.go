package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	out := DoTheThing1("input.txt")
	fmt.Printf("output is: %v\n", out)
}


func DoTheThing1(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed to open file")
	}
	scanner := bufio.NewScanner(file)

	grandTotal := 0
	for scanner.Scan() {
		rowtext := scanner.Text()

		cardInfo := strings.Split(rowtext, ": ")

		if len(cardInfo) != 2{
			panic("failed to split on :")
		}

		cardTitle := cardInfo[0]

		titleDetails := strings.Split(cardTitle, " ")
		cardNumber := titleDetails[1]

		cardDetails := strings.Split(cardInfo[1], "| ")
		winningNumbers := strings.Fields(cardDetails[0])
		gameNumbers := strings.Fields(cardDetails[1])

		cardTotal := 0
		GameNumbers:
		for _, number := range gameNumbers {
			fmt.Printf("game: %v number: %v\n", cardNumber, number)
			for _, winningNumber := range winningNumbers {
				if number == winningNumber {
					fmt.Printf("number %v is a winning number, ", number)
					if cardTotal == 0 {
						cardTotal = 1
					} else {
						cardTotal = cardTotal * 2
					}
					continue GameNumbers;
				}
			}
		}
		fmt.Printf("total for game %v is %v\n",  cardNumber, cardTotal)
		grandTotal = grandTotal + cardTotal
	}

	return grandTotal 
}

func DoTheThing2 (filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed to open file")
	}
	scanner := bufio.NewScanner(file)

	grandTotal := 0
	for scanner.Scan() {
		rowtext := scanner.Text()

		cardInfo := strings.Split(rowtext, ": ")

		if len(cardInfo) != 2{
			panic("failed to split on :")
		}

		// cardTitle := cardInfo[0]

		// titleDetails := strings.Split(cardTitle, " ")
		// cardNumber := titleDetails[1]

		cardDetails := strings.Split(cardInfo[1], "| ")
		winningNumbers := strings.Fields(cardDetails[0])
		gameNumbers := strings.Fields(cardDetails[1])

	
		GameNumbers:
		for _, number := range gameNumbers {
			for _, winningNumber := range winningNumbers {
				if number == winningNumber {
					fmt.Printf("number %v is a winning number, ", number)
					if cardTotal == 0 {
						cardTotal = 1
					} else {
						cardTotal = cardTotal * 2
					}
					continue GameNumbers;
				}
			}
		}
		grandTotal = grandTotal + cardTotal
	}
	return grandTotal 
}


