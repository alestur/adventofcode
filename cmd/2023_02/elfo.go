package main

import (
	"fmt"
	"log/slog"
	"math"
	"regexp"
	"strconv"

	"githib.com/alestur/adventofcode/pkg/common"
)

const (
	red   = "red"
	green = "green"
	blue  = "blue"
)

var color = regexp.MustCompile(fmt.Sprintf(`(%s|%s|%s)`, red, green, blue))
var num = regexp.MustCompile(`\d+`)
var setting = map[string]int{
	red:   12,
	green: 13,
	blue:  14,
}

var lineGame = regexp.MustCompile(`Game\s*([0-9]+):\s*`)
var gameSetsSep = regexp.MustCompile(`\s*;\s*`)
var setColorSep = regexp.MustCompile(`\s*,\s*`)

func main() {
	lines, err := common.ReadLines()
	if err != nil {
		slog.Error("file error", err)
		return
	}

	var resultGames int
	var resultPower float64
	for _, b := range lines {
		matchGame := lineGame.Find(b)
		game := string(b[len(matchGame):])
		gameNo, err := strconv.Atoi(string(num.Find(matchGame)))
		if err != nil {
			slog.Error("could not extract the game's number", err)
			return
		}

		gamePossible := true
		mins := map[string]float64{
			red:   0,
			green: 0,
			blue:  0,
		}
		sets := gameSetsSep.Split(game, -1)
		for _, s := range sets {
			colors := setColorSep.Split(s, -1)
			for _, c := range colors {
				cubeNo, err := strconv.Atoi(string(num.Find([]byte(c))))
				colorName := string(color.Find([]byte(c)))
				if err != nil {
					slog.Error("could not extract cube number", err)
					return
				}

				if cubeNo > setting[colorName] {
					gamePossible = false
				}

				mins[colorName] = math.Max(mins[colorName], float64(cubeNo))
			}
		}

		if gamePossible {
			resultGames += gameNo
		}

		resultPower += mins[red] * mins[green] * mins[blue]
	}

	fmt.Println("The sum of the possible games's numbers:", resultGames)
	fmt.Println("The sum of all the games' powers:", resultPower)
}
