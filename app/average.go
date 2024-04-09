package app

import (
	"fmt"
	"github.com/ediallocyf/math-skills/util"
	"math"
	"strconv"
	"strings"
)

func Average(data string) int {
	data, _ = util.ReadFileToStr("data.txt")
	var sum float64
	var count int

	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue // Skip empty lines
		}
		nums := strings.Fields(line) // Split line into individual numbers
		for _, numStr := range nums {
			num, err := strconv.ParseFloat(numStr, 64)
			if err != nil {
				fmt.Printf("Error parsing number: %v\n", err)
				continue // Skip invalid numbers
			}
			sum += num
			count++
		}
	}

	if count == 0 {
		return 0 // Avoid division by zero
	}

	average := sum / float64(count)
	return int(math.Round(average))
}
