package jz

import "testing"

var tests = []struct {
	test []int
	want int
}{
	{
		test: []int{2, 3, 4, 1, 1, 1, 2, 2},
		want: 1,
	},
	{
		test: []int{1, 1, 0, 1, 1, 1, 1, 1},
		want: 0,
	}, {
		test: []int{0, 0, 0, 0, 0},
		want: 0,
	}, {
		test: []int{2, 2, 3, 4, 1, 1},
		want: 1,
	}, {
		test: []int{},
		want: 0,
	},
	{
		test: []int{6501, 6828, 6963, 7036, 7422, 7674, 8146, 8468, 8704, 8717, 9170, 9359, 9719, 9895, 9896, 9913, 9962, 154, 293, 334, 492, 1323, 1479, 1539, 1727, 1870, 1943, 2383, 2392, 2996, 3282, 3812, 3903, 4465, 4605, 4665, 4772, 4828, 5142, 5437, 5448, 5668, 5706, 5725, 6300, 6335},
		want: 154,
	},
}

func TestMinNumberInRotatedArray(t *testing.T) {
	for _, value := range tests {
		if result := minNumberInRotatedArray(value.test); result != value.want {
			t.Errorf("minNumberInRotatedArray failed, test: %v, result: %d, want: %d", value.test, result, value.want)
		}
	}
}

func BenchmarkMinNumberInRotatedArray(b *testing.B) {
	length := len(tests)

	for i := 0; i < b.N; i++ {
		minNumberInRotatedArray(tests[i%length].test)
	}
}
