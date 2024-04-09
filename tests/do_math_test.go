package tests

import (
	"github.com/ediallocyf/math-skills/app"
	"testing"
)

// ======================================================
func TestDoMath(t *testing.T) {
	//t.Skip()
	// ------------
	t.Run("1-return ", func(t *testing.T) {
		//----Set up----
		a := 20
		b := 5
		// -------ACT--------
		outNum := app.DoMath(a, b)
		// ------expectation------
		expNum := 4
		assertDoMath(t, outNum, expNum)
	})
}

// =============assertion function ============
func assertDoMath(t testing.TB, outNum, expNum int) {
	t.Helper()
	if outNum != expNum {
		t.Errorf("\n>>>>>receive:\n%v-expect:\n%v\n", outNum, expNum)
	}
}

// ===============================================
