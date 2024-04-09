package tests

import (
	"github.com/ediallocyf/math-skills/app"
	"github.com/ediallocyf/math-skills/util"
	"log"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// ======================================================
func TestMedian(t *testing.T) {
	//t.Skip()
	// ------------
	t.Run("1-return ", func(t *testing.T) {
		//----Set up----
		dataArr := make([]float64, 1000)
		outNum := 0.0
		nums, err := util.ReadFileToStr("nums.txt")
		lines := strings.Split(string(nums), "\n")
		if err != nil {
			t.Error(err)
		}
		for _, line := range lines {
			line = strings.TrimSpace(line)
			floatNum, err := strconv.ParseFloat(line, 64)
			if err != nil {
				log.Fatal(err)
			}
			dataArr = append(dataArr, floatNum)
			sort.Float64s(dataArr)
			n := len(dataArr)
			if n%2 == 0 {
				outNum = app.Median(string(dataArr[n/2-1]))
			}
		}
		// -------ACT--------

		// ------expectation------
		expNum := 910.0
		assertMedian(t, outNum, expNum)
	})
}

// =============assertion function ============
func assertMedian(t testing.TB, outNum float64, expNum float64) {
	t.Helper()
	if outNum != expNum {
		t.Errorf("\n>>>>>receive:\n%v-expect:\n%v\n", outNum, expNum)
	}
}

// ===============================================
