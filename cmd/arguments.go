package cmd

import (
	"fmt"
	"strconv"
)

const (
	MAX_COUNT     = 1_000
	DEFAULT_COUNT = 1
)

func getCountFromArgs(args []string) (int, error) {
	if len(args) == 0 {
		return DEFAULT_COUNT, nil
	}

	count, err := strconv.Atoi(args[0])
	if err != nil || count <= 0 {
		return 0, fmt.Errorf("Argument Error: The count must be a positive integer.")
	}

	if count > MAX_COUNT {
		return 0, fmt.Errorf("Argument Error: The count of UUIDs cannot exceed %d.", MAX_COUNT)
	}

	return count, nil
}
