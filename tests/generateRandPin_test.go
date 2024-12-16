package tests

import (
	"testing"

	"dkl.dklsa.mailer/iternal"
)

type TestPinTableData struct {
	expectedLength int
	minInRange     int
	maxInRange     int
}

func newTestPinTableData(n int, min int, max int) TestPinTableData {
	return TestPinTableData{expectedLength: n, minInRange: min, maxInRange: max}
}

func TestGenerateRandPin(t *testing.T) {
	testTable := []TestPinTableData{
		newTestPinTableData(1, 0, 9),
		newTestPinTableData(2, 10, 99),
		newTestPinTableData(3, 100, 999),
		newTestPinTableData(4, 1000, 9999),
	}

	//act
	for _, testData := range testTable {

		//assert
		result := iternal.GenerateRandPin(testData.expectedLength)

		//arrange
		if !(result >= testData.minInRange && result <= testData.maxInRange) {
			t.Errorf(
				"Create pin failed: expected %v, got %v, min = %v, max = %v",
				result,
				testData.expectedLength,
				testData.minInRange,
				testData.maxInRange,
			)
		}
	}
}
