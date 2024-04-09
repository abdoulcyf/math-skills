package app

import (
	"github.com/ediallocyf/math-skills/util"
	"math"
	"strconv"
	"strings"
)

func Variance(data string) int {
	data, _ = util.ReadFileToStr("data.txt")
	var dataArr []float64
	avg := Average(data)

	// Convert data string to float64 slice
	for _, line := range strings.Split(data, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue // Skip empty lines
		}
		nums := strings.Fields(line) // Split line into individual numbers
		for _, numStr := range nums {
			floatNum, err := strconv.ParseFloat(numStr, 64)
			if err != nil {
				panic(err) // Handle parsing error
			}
			dataArr = append(dataArr, floatNum)
		}
	}

	// Calculate sum of squared differences
	sumSquaredDiff := 0.0
	for _, num := range dataArr {
		diff := num - float64(avg)
		sumSquaredDiff += math.Pow(diff, 2)
	}

	// Calculate variance
	variance := math.Round(sumSquaredDiff / float64(len(dataArr)))

	return int(variance)
}
