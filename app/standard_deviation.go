package app

import (
	"github.com/ediallocyf/math-skills/util"
	"math"
)

func StandardDeviation(data string) int {
	data, _ = util.ReadFileToStr("data.txt")
	variance := Variance(data)
	return int(math.Round(math.Sqrt(float64(variance))))
}
