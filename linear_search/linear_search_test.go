package linear_search

import "testing"

type linearSearchTest struct {
	arr                      []int64
	arrLen, target, expected int64
}

var arr = []int64{0, 0, 4, 7, 10, 23, 34, 40, 55, 68, 77, 90}
var arrLen = int64(len(arr) - 1)

var linearSearchTests = []linearSearchTest{
	{arr, arrLen, 10, 4},
	{arr, arrLen, 69, -1},
}

func TestLinearSearch(t *testing.T) {
	for _, test := range linearSearchTests {
		if output := LinearSearch(test.arr, test.arrLen, test.target); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
