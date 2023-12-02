package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./1.input.txt")
	if err != nil {
		panic(fmt.Errorf("file not found"))
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var total int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		var secondDigit int
		var lineNumber int
		for _, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				continue
			}
			if lineNumber == 0 {
				lineNumber = 10 * num
			}
			secondDigit = num
		}
		lineNumber += secondDigit
		fmt.Printf("line total is: %v", lineNumber)
		total += lineNumber
	}
	fmt.Printf("total is: %v", total)

}
