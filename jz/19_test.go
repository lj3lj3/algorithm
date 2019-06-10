package jz

import "testing"

func TestRegMatch(t *testing.T) {
	tests := []struct {
		test []rune
		reg  []rune
		want bool
	}{
		{
			test: []rune(""),
			reg:  []rune("."),
			want: false,
		},
		{
			test: []rune("aaa"),
			reg:  []rune("a.a"),
			want: true,
		},
		{
			test: []rune("aaa"),
			reg:  []rune("ab*ac*a"),
			want: true,
		},
		{
			test: []rune("aaa"),
			reg:  []rune("aa.a"),
			want: false,
		},
		{
			test: []rune("aaa"),
			reg:  []rune("ab*a"),
			want: false,
		},
		{
			test: []rune("aa"),
			reg:  []rune("a"),
			want: false,
		},
		{
			test: []rune("aa"),
			reg:  []rune("a*"),
			want: true,
		},
		{
			test: []rune("aab"),
			reg:  []rune("c*a*b"),
			want: true,
		},
		{
			test: []rune("mississippi"),
			reg:  []rune("mis*is*p*."),
			want: false,
		},
	}

	for _, value := range tests {
		if result := RegMatch(value.test, value.reg); result != value.want {
			t.Errorf("RegMatch failed, test: %s, reg: %s, results: %t, want: %t", string(value.test), string(value.reg), result, value.want)
		}
	}
}
