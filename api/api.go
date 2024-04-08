package api

import (
	"fmt"
	"github.com/ediallocyf/math-skills/util"
	"math"
	"strconv"
	"strings"
)

func DoMath(a, b int) int {
	return a / b
}

func Average() float64 {
	data, err := util.ReadFileToStr("data.txt")
	if err != nil {
		fmt.Println(err)
	}
	average := 0.0
	sum := 0.0
	dataArr := make([]string, 0)
	for _, line := range strings.Split(string(data), "\n") {
		dataArr = append(dataArr, line)
		for _, num := range strings.Split(line, " ") {
			flotNum, _ := strconv.ParseFloat(num, 64)
			sum += flotNum
			average = math.Round(sum / float64(len(dataArr)))
		}
	}

	return average
}

func Median() float64 {

}
