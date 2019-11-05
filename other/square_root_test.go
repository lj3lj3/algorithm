package other

import "testing"

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
	}

	for _, tt := range tests {
		// Square newton
		t.Run(tt.name, func(t *testing.T) {
			if got := squareRootNewton(tt.args.number); got-tt.want > 1e-6 {
				t.Errorf("squareRootNewton() = %v, want %v", got, tt.want)
			}
		})
	}
}
