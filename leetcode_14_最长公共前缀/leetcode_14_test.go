package leetcode_14_最长公共前缀

import (
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{[]string{"flowerfda", "flowerceq", "flowerces"}},
			"flower",
		},
		{
			"2",
			args{[]string{"", "flowerceq", "flowerces"}},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
