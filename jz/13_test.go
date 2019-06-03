package jz

import "testing"

func TestMovingCount(t *testing.T) {
	tests := []struct {
		test int
		rows int
		cols int
		want int
	}{
		{
			test: 5,
			rows: 10,
			cols: 10,
			want: 21,
		}, {
			test: 15,
			rows: 20,
			cols: 20,
			want: 359,
		},
	}

	for _, value := range tests {
		if result := MovingCount(value.test, value.rows, value.cols); result != value.want {
			t.Errorf("MovingCount failed, test: %v, rows: %d, cols: %d, result: %d, want: %d", value.test, value.rows, value.cols, result, value.want)
		}
	}

}
