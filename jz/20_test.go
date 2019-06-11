package jz

import (
	"sync"
	"testing"
)

func TestIsNumeric(t *testing.T) {
	tests := []struct {
		test []rune
		want bool
	}{
		{
			test: []rune(""),
			want: false,
		},
		{
			test: []rune("2."),
			want: true,
		},
		{
			test: []rune("+.9"),
			want: true,
		},
		{
			test: []rune("46.e3"),
			want: true,
		},
		{
			test: []rune("+."),
			want: false,
		},
		{
			test: []rune("2. "),
			want: true,
		},
		{
			test: []rune("aaa"),
			want: false,
		},
		{
			test: []rune("+100"),
			want: true,
		},
		{
			test: []rune("5e2"),
			want: true,
		},
		{
			test: []rune("-123"),
			want: true,
		},
		{
			test: []rune("3.1416"),
			want: true,
		},
		{
			test: []rune("-1E-16"),
			want: true,
		},
		{
			test: []rune("12e"),
			want: false,
		},
		{
			test: []rune("1a3.14"),
			want: false,
		},
		{
			test: []rune("1.2.3"),
			want: false,
		},
		{
			test: []rune("+-5"),
			want: false,
		},
		{
			test: []rune("12e+4.3"),
			want: false,
		},
	}

	// init goroutine
	wg := &sync.WaitGroup{}

	for _, valueTmp := range tests {
		wg.Add(1)

		value := valueTmp
		go func() {
			if result := IsNumeric(value.test); result != value.want {
				t.Errorf("IsNumeric failed, test: %s, results: %t, want: %t", string(value.test), result, value.want)
			}

			wg.Done()
		}()
	}

	wg.Wait()
}
