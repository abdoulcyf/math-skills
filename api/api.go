package api

import (
	"fmt"
	"github.com/ediallocyf/math-skills/util"
	"math"
	"sort"
	"strconv"
	"strings"
)

func DoMath(a, b int) int {
	return a / b
}
func Average(data string) float64 {
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
	return math.Round(average)
}

func Median(data string) float64 {
	data, _ = util.ReadFileToStr("data.txt")
	var dataArr []float64
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

	sort.Float64s(dataArr) // Sort the data array after adding all numbers

	n := len(dataArr)
	if n%2 == 0 {
		// If even number of elements, take the average of the middle two elements
		return (dataArr[n/2] + dataArr[n/2-1]) / 2
	} else {
		// If odd number of elements, take the middle element
		return dataArr[n/2]
	}
}

func Variance(data string) float64 {
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
		diff := num - avg
		sumSquaredDiff += math.Pow(diff, 2)
	}

	// Calculate variance
	variance := math.Round(sumSquaredDiff / float64(len(dataArr)))

	return variance
}

func StandardDeviation(data string) float64 {
	data, _ = util.ReadFileToStr("data.txt")
	variance := Variance(data)
	return math.Round(math.Sqrt(variance))
}
