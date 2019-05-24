package jz

import "testing"

func TestJumpFloorII(t *testing.T) {
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
			test: 3,
			want: 4,
		}, {
			test: 11,
			want: 1024,
		},
	}

	for _, value := range tests {
		if result := jumpFloorII(value.test); result != value.want {
			t.Errorf("JumpFloorII failed, test: %v, result: %d, want: %d", value.test, result, value.want)
		}
	}
}
