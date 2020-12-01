package main

import (
	"os"
	"bufio"
	"fmt"
	"time"
	"strings"
	"strconv"
)

// solution return answers for level 1 and 2
func solution(puzzle string) (interface{}, interface{}) {
	inp := make([]int, 0)
	for _, a := range strings.Fields(puzzle) {
		ai, _ := strconv.Atoi(a)
		inp = append(inp, ai)
	}

	lvl1 := 0
	for i, a := range inp {
		for j, b := range inp {
			if i == j {
				continue
			}
			if a + b == 2020 {
				lvl1 = a*b
				break
			}
		}
		if lvl1 != 0 {
			break
		}
	}

	lvl2 := 0
	for i, a := range inp {
		for j, b := range inp {
			for k, c := range inp {
				if i == j || i == k || j == k {
					continue
				}
				if a + b + c == 2020 {
					lvl2 = a*b*c
					break
				}
			}
			if lvl2 != 0 {
				break
			}
		}
		if lvl2 != 0 {
			break
		}
	}

	return lvl1, lvl2
}

func main() {


	test_outcome := testcases.Execute(solution)
	if test_outcome {
		// If the tests passed, we check our solution to the puzzle!

		puzzle_input := make([]string, 0)

 		// read from stdin
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			text := scanner.Text()
			puzzle_input = append(puzzle_input, text)
		}
		if scanner.Err() != nil {
			fmt.Println("Error reading from stdin: ",scanner.Err())
		}

		start := time.Now()
		lvl1, lvl2 := solution(strings.Join(puzzle_input[:], "\n"))
		elapsed := time.Since(start)

		fmt.Printf("Level 1: %v\nLevel 2: %v\n", lvl1, lvl2)
		fmt.Printf("Program took %s\n", elapsed)

	} else {
		// to be able to automate
		os.Exit(1)
	}

}