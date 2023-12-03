package main

import (
	"fmt"
	"log/slog"
	"regexp"

	"githib.com/alestur/adventofcode/pkg/common"
	"githib.com/alestur/adventofcode/pkg/trebuchet"
)

var digitPattern = regexp.MustCompile("\\d")
var digitAlphaPattern = regexp.MustCompile("(\\d|one|two|three|four|five|six|seven|eight|nine)")
var digitAlphaReversed = regexp.MustCompile("(\\d|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin)")

func main() {
	lines, err := common.ReadLines()
	if err != nil {
		slog.Error("file error", err)
	}

	var count int
	var countAlpha int
	for _, b := range lines {
		digit := digitPattern.Find(b)
		if len(digit) == 0 {
			slog.Error(fmt.Sprintf("line does not contain any digits: %q", b))
			return
		}
		num, err := trebuchet.ReadInt(digit)
		if err != nil {
			slog.Error("could not read digit as an integer", err)
			return
		}
		count += 10 * num

		digit = digitAlphaPattern.Find(b)
		if len(digit) == 0 {
			slog.Error(fmt.Sprintf("line does not contain any digits: %q", b))
			return
		}
		num, err = trebuchet.ReadInt(digit)
		if err != nil {
			slog.Error("could not read digit as an integer", err)
			return
		}
		countAlpha += 10 * num

		sr := make([]byte, len(b))
		for i := 0; i < len(b); i++ {
			sr[i] += b[len(b)-i-1]
		}

		digit = digitPattern.Find(sr)
		if len(digit) == 0 {
			slog.Error(fmt.Sprintf("line does not contain any digits backwards: %q", b))
			return
		}
		num, err = trebuchet.ReadInt(digit)
		if err != nil {
			slog.Error("could not read digit as an integer", err)
			return
		}
		count += num

		digit = digitAlphaReversed.Find(sr)
		if len(digit) == 0 {
			slog.Error(fmt.Sprintf("line does not contain any digits backwards: %q", b))
			return
		}
		num, err = trebuchet.ReadInt(digit)
		if err != nil {
			slog.Error("could not read digit as an integer", err)
			return
		}
		countAlpha += num
	}

	fmt.Println("The sum of all of the calibration values:", count)
	fmt.Println("The sum of all of the calibration values (with digits expressed as a word):", countAlpha)
}
