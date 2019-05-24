package jz

import "testing"

func TestJumpFloor(t *testing.T) {
	tests := []struct {
		test uint
		want uint
	}{
		{
			test: 0,
			want: 0,
		},
		{
			test: 1,
			want: 1,
		}, {
			test: 2,
			want: 2,
		}, {
			test: 8,
			want: 34,
		}, {
			test: 19,
			want: 6765,
		},
	}

	for _, value := range tests {
		if result := JumpFloor(value.test); result != value.want {
			t.Errorf("JumpFloor failed, test: %v, result: %d, want: %d", value.test, result, value.want)
		}
	}
}
