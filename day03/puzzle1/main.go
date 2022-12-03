package main

import (
	"bufio"
	"os"
	"strings"
)

// enum of priorities 1-52 for a-Z
const (
	a = iota + 1
	b
	c
	d
	e
	f
	g
	h
	i
	j
	k
	l
	m
	n
	o
	p
	q
	r
	s
	t
	u
	v
	w
	x
	y
	z
	A
	B
	C
	D
	E
	F
	G
	H
	I
	J
	K
	L
	M
	N
	O
	P
	Q
	R
	S
	T
	U
	V
	W
	X
	Y
	Z
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
		ln := strings.Split(scanner.Text(), " ")
		games = append(games, ln)
	}

	// initialise a score counter
	score := 0
}
