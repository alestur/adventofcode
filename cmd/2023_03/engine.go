package main

import (
	"fmt"
	"log/slog"

	"githib.com/alestur/adventofcode/pkg/common"
)

const gear = '*'

type part struct {
	i   int
	j   int
	typ byte
	num []int
}

func scanColumn(i int, j int, parts map[int]map[int]part) (*part, bool) {
	if p, ok := parts[i][j]; ok {
		return &p, true
	}

	if p, ok := parts[i-1][j]; ok {
		return &p, true
	}

	if p, ok := parts[i+1][j]; ok {
		return &p, true
	}

	return nil, false
}

func findPart(i int, j int, num int, parts map[int]map[int]part) (*part, bool) {
	var logN int
	for order := 1; num/order > 0; order *= 10 {
		logN += 1
	}

	for x := j - 1; x-j <= logN; x++ {
		if p, ok := scanColumn(i, x, parts); ok {
			return p, ok
		}
	}

	return nil, false
}

func readNumber(b []byte) int {
	var n int
	for _, ch := range b {
		i := int(ch) - 48
		if i < 0 || i > 9 {
			return n
		}

		n = 10*n + i
	}

	return n
}

func main() {
	lines, err := common.ReadLines()
	if err != nil {
		slog.Error("file error", err)
		return
	}
	numbers := make(map[int]map[int]int)
	parts := make(map[int]map[int]part)

	for i, ln := range lines {
		numbers[i] = make(map[int]int)
		parts[i] = make(map[int]part)

		inNumber := false

		for j, ch := range ln {
			switch ch {
			case '.':
				inNumber = false
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				if !inNumber {
					numbers[i][j] = readNumber(ln[j:])
				}
				inNumber = true
			default:
				inNumber = false
				parts[i][j] = part{
					i:   i,
					j:   j,
					typ: ch,
					num: []int{},
				}
			}
		}
	}

	var partSum int
	for i, row := range numbers {
		for j, n := range row {
			p, ok := findPart(i, j, n, parts)
			if !ok {
				continue
			}

			partSum += n

			if p.typ == gear {
				p.num = append(p.num, n)
				parts[p.i][p.j] = *p
			}
		}
	}

	var gearRatioSum int
	for _, row := range parts {
		for _, p := range row {
			if p.typ == gear && len(p.num) == 2 {
				gearRatioSum += p.num[0] * p.num[1]
			}
		}
	}

	fmt.Println("The sum of all numbers adjent to a part:", partSum)
	fmt.Println("The sum of all gear ratios:", gearRatioSum)
}
