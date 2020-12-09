package execution

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	tc "github.com/IAmBullsaw/AOC-2020/pkg/testcase"
)

func Run(solution_lvl1 func(string, map[string]int) int,
	lvl1Parameters map[string]int,
	solution_lvl2 func(string, map[string]int) int,
	lvl2Parameters map[string]int,
	testcases tc.TestCases) {
	test_outcome := testcases.Execute(solution_lvl1, solution_lvl2)
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
			fmt.Println("Error reading from stdin: ", scanner.Err())
		}

		puzzle := strings.Join(puzzle_input[:], "\n")

		fmt.Println("Solving puzzle...")
		start := time.Now()
		lvl1 := solution_lvl1(puzzle, lvl1Parameters)
		elapsed := time.Since(start)
		fmt.Printf("  Level 1 took %s\n", elapsed)

		start = time.Now()
		lvl2 := solution_lvl2(puzzle, lvl2Parameters)
		elapsed = time.Since(start)
		fmt.Printf("  Level 2 took %s\n", elapsed)

		fmt.Printf("\nAnswer\n  Level 1: %v\n  Level 2: %v\n", lvl1, lvl2)

	} else {
		// to be able to automate
		os.Exit(1)
	}
}
