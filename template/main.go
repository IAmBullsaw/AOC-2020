package main

import (
	"os"
	"bufio"
	"fmt"
	"time"
	"strings"
)

// solution return answers for level 1 and 2
func solution(puzzle string) (interface{}, interface{}) {
	return 1, 2
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