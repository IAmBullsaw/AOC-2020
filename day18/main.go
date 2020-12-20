package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/IAmBullsaw/AOC-2020/pkg/execution"
)

func evaluate(lhs int, op rune, rhs int) int {
	if op == '+' {
		return lhs + rhs
	} else if op == '*' {
		return lhs * rhs
	} else if op == '-' {
		return lhs - rhs
	} else if op == '/' {
		return lhs / rhs
	}
	panic(fmt.Sprintf("%d %s %d", lhs, string(op), rhs))
}

func pop(l *[]int) int {
	if len(*l) == 0 {
		panic("Popping empty list!")
	}
	el := (*l)[len(*l)-1]
	*l = (*l)[:len(*l)-1]
	return el
}

func popRune(l *[]rune) rune {
	el := (*l)[len(*l)-1]
	*l = (*l)[:len(*l)-1]
	return el
}

func precOper(char rune) int {
	return precByte(byte(char))
}

func precByte(char byte) int {
	if char == byte('+') || char == byte('*') {
		return 1
	} else {
		return 0
	}
}

func evaluateExpr(expr string) int {
	operands := []rune{}
	values := []int{}
	i := 0
	for i < len(expr) {
		if expr[i] == ' ' || unicode.IsSpace(rune(expr[i])) {
			i++
			continue
		} else if expr[i] == '(' {
			operands = append(operands, '(')
			i++
			continue
		} else if unicode.IsDigit(rune(expr[i])) {
			value := 0
			for i < len(expr) && unicode.IsDigit(rune(expr[i])) {
				v, _ := strconv.Atoi(string(expr[i]))
				value = (value * 10) + v
				i++
			}
			values = append(values, value)
			i--
		} else if expr[i] == ')' {
			for len(operands) != 0 && operands[len(operands)-1] != '(' {
				rhs := pop(&values)
				lhs := pop(&values)
				op := popRune(&operands)
				values = append(values, evaluate(lhs, op, rhs))
			}
			popRune(&operands)
		} else {
			for len(operands) != 0 &&
				precOper(operands[len(operands)-1]) >= precByte(expr[i]) {
				rhs := pop(&values)
				lhs := pop(&values)
				op := popRune(&operands)
				values = append(values, evaluate(lhs, op, rhs))
			}
			operands = append(operands, rune(expr[i]))
		}
		i++
	}
	for len(operands) != 0 {
		rhs := pop(&values)
		lhs := pop(&values)
		op := popRune(&operands)
		values = append(values, evaluate(lhs, op, rhs))
	}
	return values[len(values)-1]
}

// solutionLvl1 return answers for level 1
func solutionLvl1(puzzle string, parameters map[string]int) (answer int) {
	for _, line := range strings.Split(puzzle, "\n") {
		if line != "" {
			answer += evaluateExpr(line)
		}
	}
	return
}

// solutionLvl2 return answers for level 2
func solutionLvl2(puzzle string, parameters map[string]int) (answer int) {
	return
}

var lvl1Parameters = map[string]int{}
var lvl2Parameters = map[string]int{}

func main() {
	execution.Run(solutionLvl1, lvl1Parameters, solutionLvl2, lvl2Parameters, testcases)
}
