package jz

import "testing"

func TestFibonacci(t *testing.T) {
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
			want: 1,
		}, {
			test: 9,
			want: 34,
		}, {
			test: 20,
			want: 6765,
		},
	}

	for _, value := range tests {
		if result := Fibonacci(value.test); result != value.want {
			t.Errorf("fibonacci failed, test: %v, result: %d, want: %d", value.test, result, value.want)
		}
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(60)
	}
}
