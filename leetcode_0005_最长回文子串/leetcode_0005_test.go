package leetcode_0005_最长回文子串

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{"abcdcba"},
			"abcdcba",
		},
		{
			"2",
			args{"abcdcbd"},
			"bcdcb",
		},
		{
			"3",
			args{"abcde"},
			"e",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
