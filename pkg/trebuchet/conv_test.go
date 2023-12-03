package trebuchet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadInt(t *testing.T) {
	tests := map[string]int{
		"0":     0,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"eno":   1,
		"owt":   2,
		"eerht": 3,
		"ruof":  4,
		"evif":  5,
		"xis":   6,
		"neves": 7,
		"thgie": 8,
		"enin":  9,
	}

	for s, expected := range tests {
		t.Run(s, func(t *testing.T) {
			n, err := ReadInt([]byte(s))
			assert.NoError(t, err)
			assert.Equal(t, expected, n)
		})
	}
}
