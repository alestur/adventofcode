package common

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
)

func ReadLines() ([][]byte, error) {
	if len(os.Args) < 2 {
		return nil, fmt.Errorf("no input file given")
	}

	input, err := os.Open(os.Args[1])
	if err != nil {
		return nil, fmt.Errorf("opening file %q: %w", os.Args[0], err)
	}
	defer func() {
		input.Close()
		slog.Info(fmt.Sprintf("file %q closed", input.Name()))
	}()
	slog.Info(fmt.Sprintf("file %q opened", input.Name()))

	fs := bufio.NewScanner(input)
	fs.Split(bufio.ScanLines)

	lines := make([][]byte, 0)
	for fs.Scan() {
		lines = append(lines, []byte(fs.Text()))
	}

	return lines, nil
}
