package main

import (
	"fmt"
	"log/slog"
	"regexp"
	"strconv"

	"githib.com/alestur/adventofcode/pkg/common"
)

type card struct {
	winning map[int]bool
	hand    []int
}

func (c card) evaluate() int {
	var score int
	for _, i := range c.hand {
		if c.winning[i] {
			score++
		}
	}

	return score
}

var num = regexp.MustCompile(`\d+`)
var lineCard = regexp.MustCompile(`Card\s*([0-9]+):\s*`)
var numSep = regexp.MustCompile(`\s+`)
var setSep = regexp.MustCompile(`\s*\|\s*`)

func rowToInt(s string) ([]int, error) {
	chunks := numSep.Split(s, -1)
	nums := make([]int, 0, len(chunks))

	for _, ch := range chunks {
		i, err := strconv.Atoi(ch)
		if err != nil {
			return nil, err
		}

		nums = append(nums, i)
	}

	return nums, nil
}

func copyCount(cards map[int]card, from int, to int) int {
	cnt := 1
	for i := from; i < to; i++ {
		card, ok := cards[i]
		if !ok {
			continue
		}

		cnt += copyCount(cards, i+1, i+card.evaluate()+1)
	}

	return cnt
}

func main() {
	lines, err := common.ReadLines()
	if err != nil {
		slog.Error("file error", err)
		return
	}

	cards := make(map[int]card, len(lines))
	for _, ln := range lines {
		matchCard := lineCard.Find(ln)
		cardNo, err := strconv.Atoi(string(num.Find(matchCard)))
		if err != nil {
			slog.Error("could not extract the card's number", err)
			return
		}

		deal := setSep.Split(string(ln[len(matchCard):]), 2)
		if len(deal) < 2 {
			slog.Error("bad deal", slog.Int("card", cardNo), err)
			return
		}

		wins, err := rowToInt(deal[0])
		if err != nil {
			slog.Error("number was not readable", deal[0], err)
			return
		}
		hand, err := rowToInt(deal[1])
		if err != nil {
			slog.Error("number was not readable", deal[1], err)
			return
		}

		cards[cardNo] = card{
			hand:    hand,
			winning: make(map[int]bool, len(wins)),
		}
		for _, w := range wins {
			cards[cardNo].winning[w] = true
		}
	}

	var result int
	for _, c := range cards {
		score := c.evaluate()
		inc := 0
		for n := 0; n < score; n++ {
			if inc == 0 {
				inc = 1
				continue
			}
			inc *= 2
		}
		result += inc
	}

	var scratchPads = copyCount(cards, 0, len(cards))

	fmt.Println("The cards are altogether worth:", result)
	fmt.Println("The total number of scratchpads including copies:", scratchPads)
}
