package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/IAmBullsaw/AOC-2020/pkg/execution"
)

func applyMask(bin, mask string) string {
	var buffer bytes.Buffer

	for i, n := range mask {
		if n != 'X' {
			buffer.WriteRune(n)
		} else {
			buffer.WriteByte(bin[i])
		}
	}
	return buffer.String()
}

func to36Bin(s string) string {
	var i int64
	i, _ = strconv.ParseInt(s, 10, 36)

	binstr := strconv.FormatInt(i, 2)
	var buffer bytes.Buffer
	padding := 36 - len(binstr)
	for padding != 0 {
		buffer.WriteRune('0')
		padding--
	}
	buffer.WriteString(binstr)
	return buffer.String()
}

// solutionLvl1 return answers for level 1
func solutionLvl1(puzzle string, parameters map[string]int) (answer int) {
	mem := map[string]string{}
	mask := ""
	for _, line := range strings.Split(puzzle, "\n") {
		if line == "" {
			continue
		}
		data := strings.Split(line, "=")
		if strings.HasPrefix(data[0], "ma") {
			mask = data[1][1:]
			continue
		}
		address := data[0][4 : len(data[0])-2]
		value := data[1][1:]

		masked := applyMask(to36Bin(value), mask)
		mem[address] = masked
	}
	for _, v := range mem {
		val, err := strconv.ParseInt(v, 2, 64) // now we decimal again
		if err != nil {
			fmt.Println("HOLY SHIT", err)
		}
		answer += int(val)
	}
	return
}

func applyAddressMask(address, mask string) string {
	var buffer bytes.Buffer
	for i, n := range mask {
		if n == 'X' || n == '1' {
			buffer.WriteRune(n)
		} else {
			buffer.WriteByte(address[i])
		}
	}

	return buffer.String()
}

func generateAdresses(addr string) []string {
	adresses := []*bytes.Buffer{&bytes.Buffer{}}
	direction := true
	for _, c := range addr {
		if c == 'X' {
			size := len(adresses) - 1
			for ; 0 <= size; size-- {
				var cp bytes.Buffer
				cp.WriteString(adresses[size].String())
				adresses = append(adresses, &cp)
			}
			for i, adress := range adresses {
				if direction {
					if i%2 == 0 {
						adress.WriteRune('0')
					} else {
						adress.WriteRune('1')
					}
				} else {
					if i%2 == 0 {
						adress.WriteRune('1')
					} else {
						adress.WriteRune('0')
					}
				}
			}
			direction = !direction
		} else {
			for _, adress := range adresses {
				adress.WriteRune(c)
			}
		}
	}
	var res []string
	for _, a := range adresses {
		res = append(res, a.String())
	}
	return res
}

// solutionLvl2 return answers for level 2
func solutionLvl2(puzzle string, parameters map[string]int) (answer int) {
	mem := map[string]string{}
	mask := ""
	for _, line := range strings.Split(puzzle, "\n") {
		if line == "" {
			continue
		}
		data := strings.Split(line, "=")
		if strings.HasPrefix(data[0], "ma") {
			mask = data[1][1:]
			continue
		}
		address := data[0][4 : len(data[0])-2]
		value := data[1][1:]

		masked := applyAddressMask(to36Bin(address), mask)
		adresses := generateAdresses(masked)
		for _, m := range adresses {
			mem[m] = value
		}
	}
	for _, v := range mem {
		val, _ := strconv.Atoi(v)
		answer += val
	}

	return
}

var lvl1Parameters = map[string]int{}
var lvl2Parameters = map[string]int{}

func main() {
	execution.Run(solutionLvl1, lvl1Parameters, solutionLvl2, lvl2Parameters, testcases)
}
