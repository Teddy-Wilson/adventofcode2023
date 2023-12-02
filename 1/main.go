package main

import ("os" 
	"fmt"
	"bufio"
"strconv")

func main(){
	result:= SolveTheThing2("input.txt")
	fmt.Printf("total is: %v", result)
}

func SolveTheThing1(filename string)int {

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
		total += lineNumber
	}
	return total
}

func SolveTheThing2(filename string)int {
	numberPatterns := map[string]int{
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	f, err := os.Open("./1.input.txt")
	if err != nil {
		panic(fmt.Errorf("file not found"))
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var total int
	for scanner.Scan() {
		line := []rune(scanner.Text())
		if len(line) == 0 {
			break
		}
		var firstDigit int
		var lastDigit int
		for i := range line {
			MatchPatternLoop:
			for key, val :=  range numberPatterns {
				runeKey := []rune(key)
				sliceLength := len(key)
				if i + sliceLength > len(line){
					continue;
				}
				subStringInLine := line[i:i+sliceLength]
				for i, rune := range subStringInLine {
					if rune != runeKey[i] {
						continue MatchPatternLoop
					}
				}
				if firstDigit == 0 {
					firstDigit = val
				}
				lastDigit = val;
			}
		}
		lineNumber := firstDigit * 10 + lastDigit
		total += lineNumber
	}
	return total
}
