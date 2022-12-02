package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// create an empty 2d array (slice)
	games := [][]string{}

	// open the input file
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// scan the input file line by line into the array
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ln := strings.Split(scanner.Text(), " ")
		games = append(games, ln)
	}

	// initialise a score counter
	score := 0

	// iterate over the array, score the elements for each game based on lose/draw/win
	for idx := range games {
		switch {
		case games[idx][0] == "A":
			switch {
			case games[idx][1] == "X":
				// lose against rock
				score += 3 // 3 points for scissors
			case games[idx][1] == "Y":
				// draw against rock
				score += 1 // 1 point for rock
				score += 3 // 3 points for a draw
			case games[idx][1] == "Z":
				// win against rock
				score += 2 // 2 points for paper
				score += 6 // 6 points for winning
			}
		case games[idx][0] == "B":
			switch {
			case games[idx][1] == "X":
				// lose against paper
				score += 1 // 1 point for rock
			case games[idx][1] == "Y":
				// draw against paper
				score += 2 // 2 points for paper
				score += 3 // 3 points for a draw
			case games[idx][1] == "Z":
				// win against paper
				score += 3 // 3 points for scissors
				score += 6 // 6 points for winning
			}
		case games[idx][0] == "C":
			switch {
			case games[idx][1] == "X":
				// lose against scissors
				score += 2 // 2 points for paper
			case games[idx][1] == "Y":
				//draw against scissors
				score += 3 // 3 points for scissors
				score += 3 // 3 points for a draw
			case games[idx][1] == "Z":
				// win against scissors
				score += 1 // 1 points for rock
				score += 6 // 6 points for winning
			}
		}
	}

	fmt.Printf("Total score is: %d\n", score)

}
