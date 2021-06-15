package leetcode_0006_Z字形变换

import (
	"testing"
)

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{
				s:       "LEETCODEISHIRING",
				numRows: 3,
			},
			"LCIRETOESIIGEDHN",
		},
		{
			"2",
			args{
				s:       "ABC",
				numRows: 1,
			},
			"ABC",
		},
		{
			"3",
			args{
				s:       "AB",
				numRows: 1,
			},
			"AB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
