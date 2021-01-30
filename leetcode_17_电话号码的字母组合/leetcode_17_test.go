package leetcode_17_电话号码的字母组合

import (
	"reflect"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"1",
			args{"23"},
			[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			"2",
			args{"234"},
			[]string{"adh", "adk", "adi", "aeh", "aek", "aei", "afh", "afk", "afi", "bdh", "bdk", "bdi", "beh", "bek", "bei", "bfh", "bfk", "bfi", "cdh", "cdk", "cdi", "ceh", "cek", "cei", "cfh", "cfk", "cfi"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterCombinations(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}
