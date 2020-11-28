package main

import (
	"fmt"

	"github.com/jblashki/aoc-filereader-go"
)

const INPUT_FILE = "./floors"

func main() {
	var day1aResult int
	var day1bResult int
	var err error

	fmt.Printf("AoC 2015 Day 1 (GO)\n")
	fmt.Printf("-------------------\n")
	day1aResult, err = day1a()
	if err != nil {
		fmt.Printf("1a: **** Error: %q ****\n", err)
	} else {
		fmt.Printf("1a: Floors              = %v\n", day1aResult)
	}
	day1bResult, err = day1b()
	if err != nil {
		fmt.Printf("1b: **** Error: %q ****\n", err)
	} else {

		fmt.Printf("1b: Floor -1 @ position = %v\n", day1bResult)
	}
}

func day1a() (int, error) {
	instructions, err := filereader.ReadIntoString(INPUT_FILE)
	if err != nil {
		return 0, err
	}

	floor := 0

	for i := 0; i < len(instructions); i++ {
		switch instructions[i] {
		case '(':
			floor++
		case ')':
			floor--
		}
	}

	return floor, nil
}

func day1b() (int, error) {
	instructions, err := filereader.ReadIntoString(INPUT_FILE)
	if err != nil {
		return 0, err
	}

	floor := 0

	for i := 0; i < len(instructions); i++ {
		switch instructions[i] {
		case '(':
			floor++
		case ')':
			floor--
		}

		if floor == -1 {
			return i + 1, nil
		}
	}

	return -1, nil
}
