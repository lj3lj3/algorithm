package jz

import "testing"

func TestPrintListFromTailToHead(t *testing.T) {
	tests := []struct {
		test *list
		want string
	}{
		{
			test: &list{
				value: "1",
				next: &list{
					value: "2",
					next: &list{
						value: "3",
					},
				},
			},
			want: "3,2,1",
		},
		{
			test: &list{
				value: "5",
				next: &list{
					value: "2",
					next: &list{
						value: "1",
					},
				},
			},
			want: "1,2,5",
		},
	}

	for _, value := range tests {
		if result := printListFromTailToHead(value.test); result != value.want {
			t.Errorf("printListFromTailToHead failed, test: %v, result: %s, want: %s", *value.test, result, value.want)
		}
	}
}
