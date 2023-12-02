package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	result := SolveTheThing2("./input.txt")
	fmt.Printf("result is: \"%v\"", result)
}

type Game struct {
	id int
	grabs []Grab
	min_blue int
	min_red int
	min_green int
}

type Grab struct {
	num_red int 
	num_green int
	num_blue int
}

func SolveTheThing2(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		panic("file not found")

	}
	reader := bufio.NewReader(f)
	scanner := bufio.NewScanner(reader)

	var games []Game
	for scanner.Scan() {
		var game Game
		line := scanner.Text()
		gameInfo := strings.Split(line, ": ")
		if len(gameInfo)!=2{
			panic("error getting game info")
		}
		gameTitle := strings.Split(gameInfo[0], " ")
		if len(gameTitle) != 2 {
			panic("error deciphering game title")
		}
		gameNumberAsString := gameTitle[1]
		gameNumber, err:= strconv.Atoi(gameNumberAsString)
		if err != nil {
			panic("error deciphering game title")
		}
		game.id = gameNumber
		grabs := strings.Split(gameInfo[1], "; ")
		for  _, grab := range grabs {
			var thisGrab Grab
			for _, balls := range strings.Split(grab, ", ") {
				ball := strings.Split(balls, " ")
				if len(ball) != 2{
					panic("error figuring out how ball quant and color")
				}
				ballColor := ball[1]
				ballQuantity, err :=  strconv.Atoi(ball[0])
				if err != nil {
					panic(err)
				}

				switch(ballColor){
				case "blue":
					if thisGrab.num_blue != 0 {
						panic("somethings gone wrong")
					}
					thisGrab.num_blue = ballQuantity
				case "red":
					if thisGrab.num_red != 0{
						panic("somethings gone wrong")
					}
					thisGrab.num_red = ballQuantity
				case "green":
					if thisGrab.num_green != 0 {
						panic("sometings gone wrong")
					}
					thisGrab.num_green = ballQuantity
				default:
					panic("unknown color: " + ballColor)
				}
			}
			game.grabs = append(game.grabs, thisGrab)
		}
		
		for _, grab := range game.grabs {
			if grab.num_blue > game.min_blue {
				game.min_blue = grab.num_blue
			}
			if grab.num_green > game.min_green {
				game.min_green = grab.num_green
			}
			if grab.num_red > game.min_red {
				game.min_red = grab.num_red
			}
		}
		games = append(games, game)
	}
	var total int
	for _, game := range games {
		power := game.min_blue * game.min_red * game.min_green
		total += power
	}
	return total
}

func SolveTheThing1(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		panic("file not found")

	}
	reader := bufio.NewReader(f)
	scanner := bufio.NewScanner(reader)

	var possibleGames []Game
	for scanner.Scan() {
		var game Game
		line := scanner.Text()
		gameInfo := strings.Split(line, ": ")
		if len(gameInfo)!=2{
			panic("error getting game info")
		}
		gameTitle := strings.Split(gameInfo[0], " ")
		if len(gameTitle) != 2 {
			panic("error deciphering game title")
		}
		gameNumberAsString := gameTitle[1]
		gameNumber, err:= strconv.Atoi(gameNumberAsString)
		if err != nil {
			panic("error deciphering game title")
		}
		game.id = gameNumber
		grabs := strings.Split(gameInfo[1], "; ")
		for  _, grab := range grabs {
			var thisGrab Grab
			for _, balls := range strings.Split(grab, ", ") {
				ball := strings.Split(balls, " ")
				if len(ball) != 2{
					panic("error figuring out how ball quant and color")
				}
				ballColor := ball[1]
				ballQuantity, err :=  strconv.Atoi(ball[0])
				if err != nil {
					panic(err)
				}

				switch(ballColor){
				case "blue":
					if thisGrab.num_blue != 0 {
						panic("somethings gone wrong")
					}
					thisGrab.num_blue = ballQuantity
				case "red":
					if thisGrab.num_red != 0{
						panic("somethings gone wrong")
					}
					thisGrab.num_red = ballQuantity
				case "green":
					if thisGrab.num_green != 0 {
						panic("sometings gone wrong")
					}
					thisGrab.num_green = ballQuantity
				default:
					panic("unknown color: " + ballColor)
				}
			}
			game.grabs = append(game.grabs, thisGrab)
		}

		// grab is impossible if any grab exceeds 12 red 13 green 14 blue
		gamePossible := true
		for _, grab := range game.grabs {
			if grab.num_blue > 14 || grab.num_red > 12 || grab.num_green > 13 {
				gamePossible = false
			}
		}
		if gamePossible {
			possibleGames = append(possibleGames, game)
		}
	}
	var total int
	for _,game := range possibleGames{
		total += game.id
	}
	return total
}
