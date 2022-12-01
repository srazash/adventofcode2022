package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var input []int

	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		switch {
		case scanner.Text() == "":
			input = append(input, 0)
		default:
			ln, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}
			input = append(input, ln)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("%d lines read in from input file.\n", len(input))
	fmt.Printf("Most cals carried by an Elf: %d\n", topCal(input))
}

func topCal(in []int) int {
	var totalCal []int

	total := 0

	for idx, ele := range in {
		if ele != 0 && idx < len(in)-1 {
			total += ele
		} else {
			totalCal = append(totalCal, total)
			total = 0
		}
	}

	sort.Ints(totalCal)

	return totalCal[len(totalCal)-1]
}
