package leetcode

import (
	"testing"
)

var wordSearchTestCases = []struct {
	test   [][]byte
	search string
	want   bool
}{
	{
		test:   [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
		search: "ABCCED",
		want:   true,
	},
}

func TestHasPath(t *testing.T) {
	for _, value := range wordSearchTestCases {
		if result := Exist(value.test, value.search); result != value.want {
			t.Errorf("word search failed, test: %v, search: %v result: %t, want: %t", value.test, value.search, result, value.want)
		}
	}
}
