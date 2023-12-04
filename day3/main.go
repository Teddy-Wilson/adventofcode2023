package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	out := DoTheThing2("input.txt")
	fmt.Printf("output is: %v\n", out)
}

type Number struct {
	value     int
	locationX int
	locationY int
	perimiterLocations []Location
}
type Location struct {
	x int
	y int
}

func DoTheThing1(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed to open file")
	}
	scanner := bufio.NewScanner(file)

	var numbers []Number
	var symbolLocations []Location
	rownum := 0
	maxX := 0
	maxY := 0
	for scanner.Scan() {
		row := scanner.Text()
		maxX = len([]rune(row))
		var currNum string
		var currNumX int
		var currNumY int
		for colnum, char := range row {
			_, err := strconv.Atoi(string(char))
			if err != nil {
				if currNum != "" {
					value, err := strconv.Atoi(currNum)
					if err != nil {
						panic("somethigns gone badly wrong")
					}
					numbers = append(numbers, Number{
						value:     value,
						locationX: currNumX,
						locationY: currNumY,
					})
					currNum = ""
					currNumX = 0
					currNumY = 0
				}
			} else {
				if currNum == "" {
					currNumX = colnum
					currNumY = rownum
				}
				currNum = currNum + string(char)
			}

			switch {
			case char == '.':
				{
					continue
				}
			case err != nil:
				{
					symbolLocations = append(symbolLocations, Location{
						x: colnum,
						y: rownum,
					})
				}
			}

		}

		if currNum != "" {
			value, err := strconv.Atoi(currNum)
			if err != nil {
				panic("somethigns gone badly wrong")
			}
			numbers = append(numbers, Number{
				value:     value,
				locationX: currNumX,
				locationY: currNumY,
			})
			currNum = ""
			currNumX = 0
			currNumY = 0
		}
		rownum = rownum + 1
	}
	maxY = rownum

	var legitParts []Number
	for _, number := range numbers {
		var perimiterLocations []Location
		numberOfCharsInValue := len([]rune(strconv.Itoa(number.value)))

		var rightSideLocation Location
		if number.locationX+numberOfCharsInValue <= maxX {
			rightSideLocation = Location{
				y: number.locationY,
				x: number.locationX + numberOfCharsInValue,
			}
		} else {
			rightSideLocation = Location{
				y: number.locationY,
				x: number.locationX,
			}
		}
		var leftSideLocation Location
		if number.locationX >= 0 {
			leftSideLocation = Location{
				y: number.locationY,
				x: number.locationX - 1,
			}
		} else {
			leftSideLocation = Location{
				y: number.locationY,
				x: number.locationX,
			}
		}

		var bottomLocations []Location
		if number.locationY+1 <= maxY {
			for x := leftSideLocation.x; x <= rightSideLocation.x; x++ {
				bottomLocations = append(bottomLocations, Location{y: number.locationY + 1, x: x})
			}
		}

		var topLocations []Location
		if number.locationY >= 0 {
			for x := leftSideLocation.x; x <= rightSideLocation.x; x++ {
				topLocations = append(topLocations, Location{y: number.locationY - 1, x: x})
			}
		}

		perimiterLocations = append(perimiterLocations, rightSideLocation, leftSideLocation)
		perimiterLocations = append(perimiterLocations, topLocations...)
		perimiterLocations = append(perimiterLocations, bottomLocations...)

		isPartLegit := false
		for _, location := range perimiterLocations {
			for _, symbolLoc := range symbolLocations {
				if symbolLoc.x == location.x && symbolLoc.y == location.y {
					isPartLegit = true
				}
			}
		}
		if isPartLegit {
			legitParts = append(legitParts, number)
		}

	}

	var total int
	for _, p := range legitParts {
		total = total + p.value
	}
	return total
}

func DoTheThing2(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed to open file")
	}
	scanner := bufio.NewScanner(file)

	var numbers []Number
	var symbolLocations []Location
	var gearLocations []Location
	rownum := 0
	maxX := 0
	maxY := 0
	for scanner.Scan() {
		row := scanner.Text()
		maxX = len([]rune(row))
		var currNum string
		var currNumX int
		var currNumY int
		for colnum, char := range row {
			_, err := strconv.Atoi(string(char))
			if err != nil {
				if currNum != "" {
					value, err := strconv.Atoi(currNum)
					if err != nil {
						panic("somethigns gone badly wrong")
					}
					numbers = append(numbers, Number{
						value:     value,
						locationX: currNumX,
						locationY: currNumY,
					})
					currNum = ""
					currNumX = 0
					currNumY = 0
				}
			} else {
				if currNum == "" {
					currNumX = colnum
					currNumY = rownum
				}
				currNum = currNum + string(char)
			}

			switch {
			case char == '*': {
				gearLocations = append(gearLocations, Location{
					x: colnum,
					y: rownum,
				})
				}
			case char == '.':
				{
					continue
				}
			case err != nil:
				{
					symbolLocations = append(symbolLocations, Location{
						x: colnum,
						y: rownum,
					})
				}
			}

		}

		if currNum != "" {
			value, err := strconv.Atoi(currNum)
			if err != nil {
				panic("somethigns gone badly wrong")
			}
			numbers = append(numbers, Number{
				value:     value,
				locationX: currNumX,
				locationY: currNumY,
			})
			currNum = ""
			currNumX = 0
			currNumY = 0
		}
		rownum = rownum + 1
	}
	maxY = rownum


	for i, number := range numbers {
		var perimiterLocations []Location
		numberOfCharsInValue := len([]rune(strconv.Itoa(number.value)))

		var rightSideLocation Location
		if number.locationX+numberOfCharsInValue <= maxX {
			rightSideLocation = Location{
				y: number.locationY,
				x: number.locationX + numberOfCharsInValue,
			}
		} else {
			rightSideLocation = Location{
				y: number.locationY,
				x: number.locationX,
			}
		}
		var leftSideLocation Location
		if number.locationX >= 0 {
			leftSideLocation = Location{
				y: number.locationY,
				x: number.locationX - 1,
			}
		} else {
			leftSideLocation = Location{
				y: number.locationY,
				x: number.locationX,
			}
		}

		var bottomLocations []Location
		if number.locationY+1 <= maxY {
			for x := leftSideLocation.x; x <= rightSideLocation.x; x++ {
				bottomLocations = append(bottomLocations, Location{y: number.locationY + 1, x: x})
			}
		}

		var topLocations []Location
		if number.locationY >= 0 {
			for x := leftSideLocation.x; x <= rightSideLocation.x; x++ {
				topLocations = append(topLocations, Location{y: number.locationY - 1, x: x})
			}
		}

		perimiterLocations = append(perimiterLocations, rightSideLocation, leftSideLocation)
		perimiterLocations = append(perimiterLocations, topLocations...)
		perimiterLocations = append(perimiterLocations, bottomLocations...)
		numbers[i].perimiterLocations = perimiterLocations;
	}

	var total int
	fmt.Printf("%v\n", len(gearLocations))
	for _, gearLoc := range gearLocations {
		var matchingNumbers []Number

		Number:
		for _, number := range numbers {
			for _, p := range number.perimiterLocations {
				if p.x == gearLoc.x && p.y == gearLoc.y {
					matchingNumbers = append(matchingNumbers, number)
					continue Number;
				}
			}
		}

		if len(matchingNumbers) == 2 {
			gearRatio := 1
			for _, match := range matchingNumbers {
				gearRatio = match.value * gearRatio
			}
			total = total + gearRatio
		}
	}
	return total
}
