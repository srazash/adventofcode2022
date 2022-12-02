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

	// scan the input file line by line in the array
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ele := strings.Split(scanner.Text(), " ")
		games = append(games, ele)
	}

	fmt.Println(len(games))
	fmt.Println(games[0][0])
	fmt.Println(games[0][1])

}
