package trebuchet

import (
	"fmt"
	"strconv"
)

func ReadInt(b []byte) (int, error) {
	if len(b) == 0 {
		return -1, fmt.Errorf("zero-length string")
	}

	var n int
	var err error
	switch string(b) {
	case "one":
		n = 1
	case "two":
		n = 2
	case "three":
		n = 3
	case "four":
		n = 4
	case "five":
		n = 5
	case "six":
		n = 6
	case "seven":
		n = 7
	case "eight":
		n = 8
	case "nine":
		n = 9
	case "eno":
		n = 1
	case "owt":
		n = 2
	case "eerht":
		n = 3
	case "ruof":
		n = 4
	case "evif":
		n = 5
	case "xis":
		n = 6
	case "neves":
		n = 7
	case "thgie":
		n = 8
	case "enin":
		n = 9
	default:
		n, err = strconv.Atoi(string(b[0]))
		if err != nil {
			return -1, nil
		}
	}

	return n, nil
}
