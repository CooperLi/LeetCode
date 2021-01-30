package leetcode_20_有效的括号

import (
	"testing"
)

func Test_isValid(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{"()"},
			want: true,
		},
		{
			name: "2",
			args: args{"([])"},
			want: true,
		},
		{
			name: "3",
			args: args{"({}[])"},
			want: true,
		},
		{
			name: "4",
			args: args{"(]"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.input); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
