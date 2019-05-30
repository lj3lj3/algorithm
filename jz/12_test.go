package jz

import "testing"

var hasPathTestCases = []struct {
	test   []rune
	rows   int
	cols   int
	search []rune
	want   bool
}{
	{
		// ABCE
		// SFCS
		// ADEE
		test:   []rune("ABCESFCSADEE"),
		rows:   3,
		cols:   4,
		search: []rune("ABCCED"),
		want:   true,
	},
	{
		test:   []rune("ABCESFCSADEE"),
		rows:   3,
		cols:   4,
		search: []rune("ABCB"),
		want:   false,
	},
	{
		//ABCEHJIG
		//SFCSLOPQ
		//ADEEMNOE
		//ADIDEJFM
		//VCEIFGGS
		test:   []rune("ABCEHJIGSFCSLOPQADEEMNOEADIDEJFMVCEIFGGS"),
		rows:   5,
		cols:   8,
		search: []rune("SGGFIECVAASABCEHJIGQEM"),
		want:   true,
	},
	{
		test:   []rune("ABCEHJIGSFCSLOPQADEEMNOEADIDEJFMVCEIFGGS"),
		rows:   5,
		cols:   8,
		search: []rune("SLHECCEIDEJFGGFIE"),
		want:   true,
	},
	{
		test:   []rune("ABCEHJIGSFCSLOPQADEEMNOEADIDEJFMVCEIFGGS"),
		rows:   5,
		cols:   8,
		search: []rune("SGGFIECVAASABCEHJIGQEMS"),
		want:   false,
	},
}

func TestHasPath(t *testing.T) {
	for _, value := range hasPathTestCases {
		if result := HasPath(value.test, value.rows, value.cols, value.search); result != value.want {
			t.Errorf("HasPath failed, test: %v, search: %v result: %t, want: %t", value.test, value.search, result, value.want)
		}
	}
}

func BenchmarkHasPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, value := range hasPathTestCases {
			if result := HasPath(value.test, value.rows, value.cols, value.search); result != value.want {
				b.Errorf("HasPath failed, test: %v, search: %v result: %t, want: %t", value.test, value.search, result, value.want)
			}
		}
	}
}
