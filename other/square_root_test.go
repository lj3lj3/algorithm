package other

import (
	"testing"
	"time"
)

func TestSquare(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			args: args{number: 1},
			want: float64(1),
		},
		{
			args: args{number: 9},
			want: float64(3),
		},
		{
			args: args{number: 5},
			want: 2.236067,
		},
		{
			name: "",
			args: args{number: 2},
			want: 1.414213,
		},
		{
			name: "",
			args: args{number: 90000000000},
			want: 300000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			for i := 0; i < 1000000; i++ {
				// Square root newton
				if got := squareRootNewton(tt.args.number); got-tt.want > 1e-1 {
					t.Errorf("squareRootNewton() = %v, want %v", got, tt.want)
				}
			}
			t.Logf("Square root newton: %v", time.Since(start))

			start = time.Now()
			for i := 0; i < 1000000; i++ {
				// Square root fast
				if got := squareRootFast(tt.args.number); got-tt.want > 1e-1 {
					t.Errorf("squareRootFast() = %v, want %v", got, tt.want)
				}
			}
			t.Logf("Square root fast: %v", time.Since(start))
		})
	}
}
