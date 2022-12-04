package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//initialize empty array
	pairs := [][]string{}

	// initialise a counter
	count := 0

	// open the input file
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// scan the input file line by line into the tmp array, removing '-'
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l1 := strings.Replace(scanner.Text(), "-", ",", -1)
		l2 := strings.Split(l1, ",")
		pairs = append(pairs, l2)
	}

	for idx := range pairs {
		zero, err := strconv.Atoi(pairs[idx][0])
		one, err := strconv.Atoi(pairs[idx][1])
		two, err := strconv.Atoi(pairs[idx][2])
		three, err := strconv.Atoi(pairs[idx][3])

		if err != nil {
			panic(err)
		}

		if zero <= two && one >= three {
			count++
			fmt.Printf("✅ First pair (%d-%d) fully contains second pair (%d-%d): ", zero, one, two, three)
			fmt.Println(pairs[idx])
		} else if two <= zero && three >= one {
			count++
			fmt.Printf("✅ Second pair (%d-%d) fully contains first pair (%d-%d): ", two, three, zero, one)
			fmt.Println(pairs[idx])
		} else {
			fmt.Printf("❌ Pairs (%d-%d),(%d-%d) do not contain each other: ", zero, one, two, three)
			fmt.Println(pairs[idx])
		}
	}

	fmt.Printf("Fully overlapping pairs: %d\n", count)

}
