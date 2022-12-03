package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// create an empty 2d array (slice)
	comps := [][]string{}

	// open the input file
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// scan the input file line by line into the array
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ln := strings.Split(scanner.Text(), "")
		comps = append(comps, ln)
	}

	// initialise a score counter
	sum := 0
	count := 0

	for idx := 0; idx < len(comps); idx += 3 {
		match := false
		tmp := []string{}
		for _, e1 := range comps[idx] {
			for _, e2 := range comps[idx+1] {
				if e1 == e2 {
					tmp = append(tmp, e1)
					for _, e3 := range tmp {
						for _, e4 := range comps[idx+2] {
							if e3 == e4 && match == false {
								match = true
								count++
								sum += getPri(e3)
							}
						}
					}
				}
			}
		}
	}

	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("(Count: %d)\n", count)

}

func getPri(s string) int {
	switch {
	case s == "a":
		return 1
	case s == "b":
		return 2
	case s == "c":
		return 3
	case s == "d":
		return 4
	case s == "e":
		return 5
	case s == "f":
		return 6
	case s == "g":
		return 7
	case s == "h":
		return 8
	case s == "i":
		return 9
	case s == "j":
		return 10
	case s == "k":
		return 11
	case s == "l":
		return 12
	case s == "m":
		return 13
	case s == "n":
		return 14
	case s == "o":
		return 15
	case s == "p":
		return 16
	case s == "q":
		return 17
	case s == "r":
		return 18
	case s == "s":
		return 19
	case s == "t":
		return 20
	case s == "u":
		return 21
	case s == "v":
		return 22
	case s == "w":
		return 23
	case s == "x":
		return 24
	case s == "y":
		return 25
	case s == "z":
		return 26
	case s == "A":
		return 27
	case s == "B":
		return 28
	case s == "C":
		return 29
	case s == "D":
		return 30
	case s == "E":
		return 31
	case s == "F":
		return 32
	case s == "G":
		return 33
	case s == "H":
		return 34
	case s == "I":
		return 35
	case s == "J":
		return 36
	case s == "K":
		return 37
	case s == "L":
		return 38
	case s == "M":
		return 39
	case s == "N":
		return 40
	case s == "O":
		return 41
	case s == "P":
		return 42
	case s == "Q":
		return 43
	case s == "R":
		return 44
	case s == "S":
		return 45
	case s == "T":
		return 46
	case s == "U":
		return 47
	case s == "V":
		return 48
	case s == "W":
		return 49
	case s == "X":
		return 50
	case s == "Y":
		return 51
	case s == "Z":
		return 52
	default:
		return 0
	}
}
