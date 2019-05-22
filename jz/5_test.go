package jz

import "testing"

func TestReplaceSpace(t *testing.T) {
	tests := []struct {
		test string
		want string
	}{
		{
			test: "good one",
			want: "good%20one",
		},
		{
			test: "good one lkjkkddd",
			want: "good%20one%20lkjkkddd",
		}, {
			test: "We Are Happy",
			want: "We%20Are%20Happy",
		}, {
			test: "WeAreHappy",
			want: "WeAreHappy",
		}, {
			test: " WeAreHappy",
			want: "%20WeAreHappy",
		},
	}

	for _, value := range tests {
		if result := replaceSpace(value.test); result != value.want {
			t.Errorf("replaceSpace failed, test: %v, result: %s, want: %s", value.test, result, value.want)
		}
	}
}
