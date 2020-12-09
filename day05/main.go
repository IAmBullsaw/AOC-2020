package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	cmp "github.com/IAmBullsaw/AOC-2020/pkg/compare"
	"github.com/IAmBullsaw/AOC-2020/pkg/execution"
	ip "github.com/IAmBullsaw/AOC-2020/pkg/inp"
)

func row(bp string) int {
	bin := bp[:len(bp)-3]
	bin = strings.ReplaceAll(bin, "F", "0")
	bin = strings.ReplaceAll(bin, "B", "1")
	dec, _ := strconv.ParseInt(bin, 2, 64)
	return int(dec)
}

func col(bp string) int {
	bin := bp[len(bp)-3:]
	bin = strings.ReplaceAll(bin, "L", "0")
	bin = strings.ReplaceAll(bin, "R", "1")
	dec, _ := strconv.ParseInt(bin, 2, 64)
	return int(dec)
}

// solutionLvl1 return answers for level 1
func solutionLvl1(puzzle string, parameters map[string]int) (answer int) {
	bps := ip.ToSlice(puzzle)

	for _, bp := range bps {
		answer = cmp.Max(answer, row(bp)*8+col(bp))
	}

	return
}

type Seat struct {
	Row, Col, Id int
	Bp           string
}

func (s Seat) Nil() bool {
	return s.Id == 0
}

func SortSeats(seats []Seat) []Seat {
	sort.Slice(seats, func(i, j int) bool {
		a := seats[i]
		b := seats[j]
		if a.Row == b.Row {
			if a.Col == b.Col {
				return a.Id < b.Id
			} else {
				return a.Col < b.Col
			}
		} else {
			return a.Row < b.Row
		}
	})
	return seats
}

func MissingSeat(a, b, c Seat) (missing Seat) {
	fmt.Println(a, b, c)

	return
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

// solutionLvl2 return answers for level 2
func solutionLvl2(puzzle string, parameters map[string]int) (answer int) {
	bps := ip.ToSlice(puzzle)

	seats := []Seat{}

	for _, bp := range bps {
		seats = append(seats, Seat{Row: row(bp), Col: col(bp), Id: row(bp)*8 + col(bp), Bp: bp})
	}
	seats = SortSeats(seats)
	seatIds := makeRange(seats[0].Id, seats[len(seats)-1].Id)
	for i, seatId := range seatIds {
		if i >= len(seats) {
			break
		}
		if seats[i].Id != seatId {
			if diff := seats[i].Id - seats[i-1].Id; diff == 2 || diff == -2 {
				answer = seatId
				break
			}
		}
	}

	return
}

var lvl1Parameters = map[string]int{}
var lvl2Parameters = map[string]int{}

func main() {
	execution.Run(solutionLvl1, lvl1Parameters, solutionLvl2, lvl2Parameters, testcases)
}
