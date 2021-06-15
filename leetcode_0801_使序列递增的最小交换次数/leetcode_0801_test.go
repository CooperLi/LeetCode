package leetcode_0801_使序列递增的最小交换次数

import (
	"testing"
)

func Test_minSwap(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				a: []int{2, 3, 2, 5, 6},
				b: []int{0, 1, 4, 4, 5},
			},
			1,
		},
		{
			"2",
			args{
				a: []int{0, 2, 5, 8, 9, 10, 12, 14, 18, 19, 20, 20, 24, 27, 28, 31, 33, 34, 36, 38},
				b: []int{1, 2, 5, 7, 8, 9, 11, 17, 15, 16, 19, 21, 28, 29, 30, 31, 33, 34, 38, 39},
			},
			2,
		},
		{
			"3",
			args{
				a: []int{0, 3, 4, 9, 10},
				b: []int{2, 3, 7, 5, 6},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSwap(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("minSwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
