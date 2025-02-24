package utils

import "testing"

func TestStringCompare(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name   string
		args   args
		expect bool
	}{
		{
			name: "Equal strings",
			args: args{
				a: "hello",
				b: "hello",
			},
			expect: true,
		},
		{
			name: "Different strings",
			args: args{
				a: "hello",
				b: "world",
			},
			expect: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringCompare(tt.args.a, tt.args.b); got != tt.expect {
				t.Errorf("StringCompare() = %v, expect %v", got, tt.expect)
			}
		})
	}
}
