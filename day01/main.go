package main

import (
	"os"
	"bufio"
	"fmt"
	"time"
	"strings"
	"strconv"
	"sort"
)


// solution_lvl1 return answers for level 1
func solution_lvl1(puzzle string) (answer interface{}) {
	inp := make([]int, 0)
	for _, a := range strings.Fields(puzzle) {
		ai, _ := strconv.Atoi(a)
		inp = append(inp, ai)
	}

	sort.Ints(inp)

	i := 0
	j := len(inp)-1
	for i<=j {
		target := 2020
		sum := inp[i]+inp[j]
		if sum == target {
			answer = inp[i]*inp[j]
			break
		} else if sum < target {
			i++
		} else {
			j--
		}
	}

	return
}

// solution _lvl2 return answers for level 2
func solution_lvl2(puzzle string) (answer interface{}) {
	inp := make([]int, 0)
	for _, a := range strings.Fields(puzzle) {
		ai, _ := strconv.Atoi(a)
		inp = append(inp, ai)
	}

	sort.Ints(inp)

	for i := 0; i < len(inp); i++ {
		j := i+1
		k := len(inp)-1
		for j<=k {
			target := 2020
			sum := inp[i]+inp[j]+inp[k]
			if sum == target {
				answer = inp[i]*inp[j]*inp[k]
				break
			} else if sum < target {
				j++
			} else {
				k--
			}
		}
		if answer != nil {
			break
		}
	}

	return
}


// solution return answers for level 1 and 2 and times each execution
func solution(puzzle string) (interface{}, interface{}) {

	fmt.Println("Solving puzzle...")
	start := time.Now()
	lvl1 := solution_lvl1(puzzle)
	elapsed := time.Since(start)
	fmt.Printf("  Level 1 took %s\n", elapsed)

	start = time.Now()
	lvl2 := solution_lvl2(puzzle)
	elapsed = time.Since(start)
	fmt.Printf("  Level 2 took %s\n", elapsed)

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

		lvl1, lvl2 := solution(strings.Join(puzzle_input[:], "\n"))
		fmt.Printf("\nAnswer\n  Level 1: %v\n  Level 2: %v\n", lvl1, lvl2)

	} else {
		// to be able to automate
		os.Exit(1)
	}

}