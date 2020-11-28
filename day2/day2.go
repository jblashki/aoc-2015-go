package main

import (
	"errors"
	"fmt"

	"github.com/jblashki/aoc-filereader-go"
)

const INPUT_FILE = "./presents"

func main() {
	var day2aResult int
	var day2bResult int
	var err error

	fmt.Printf("AoC 2015 Day 2 (GO)\n")
	fmt.Printf("-------------------\n")
	day2aResult, err = day2a()
	if err != nil {
		fmt.Printf("2a: **** Error: %q ****\n", err)
	} else {
		fmt.Printf("2a: Wrapping paper = %v square feet\n", day2aResult)
	}
	day2bResult, err = day2b()
	if err != nil {
		fmt.Printf("2b: **** Error: %q ****\n", err)
	} else {
		fmt.Printf("2b: Ribbon         = %v feet\n", day2bResult)
	}
}

func day2a() (int, error) {
	lines, err := filereader.ReadLines(INPUT_FILE)
	if err != nil {
		return 0, err
	}

	total := 0
	for i := 0; i < len(lines); i++ {
		l := 0
		w := 0
		h := 0

		_, err := fmt.Sscanf(lines[i], "%dx%dx%d", &l, &w, &h)
		if err != nil {
			errormsg := fmt.Sprintf("Failed scan: '%v' %v", lines[i], err)
			return 0, errors.New(errormsg)
		}

		side1 := l * w
		side2 := w * h
		side3 := h * l

		smallest_side := side1

		if side2 < smallest_side {
			smallest_side = side2
		}

		if side3 < smallest_side {
			smallest_side = side3
		}

		paper := 2*side1 + 2*side2 + 2*side3 + smallest_side

		total += paper
	}

	return total, nil
}

func day2b() (int, error) {
	lines, err := filereader.ReadLines(INPUT_FILE)
	if err != nil {
		return 0, err
	}

	total := 0
	for i := 0; i < len(lines); i++ {
		l := 0
		w := 0
		h := 0

		_, err := fmt.Sscanf(lines[i], "%dx%dx%d", &l, &w, &h)
		if err != nil {
			errormsg := fmt.Sprintf("Failed scan: '%v' %v", lines[i], err)
			return 0, errors.New(errormsg)
		}

		side1 := 2*l + 2*w
		side2 := 2*w + 2*h
		side3 := 2*h + 2*l

		smallest_side := side1

		if side2 < smallest_side {
			smallest_side = side2
		}

		if side3 < smallest_side {
			smallest_side = side3
		}

		ribon := smallest_side + (l * w * h)

		total += ribon
	}

	return total, nil
}
