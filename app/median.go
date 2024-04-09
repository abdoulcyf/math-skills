package app

import (
	"github.com/ediallocyf/math-skills/util"
	"sort"
	"strconv"
	"strings"
)

func Median(data string) int {
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
		return int((dataArr[n/2] + dataArr[n/2-1]) / 2)
	} else {
		// If odd number of elements, take the middle element
		return int(dataArr[n/2])
	}
}
