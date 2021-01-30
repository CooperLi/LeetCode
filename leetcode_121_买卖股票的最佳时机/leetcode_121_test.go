package leetcode_121_买卖股票的最佳时机

import (
	"testing"
)

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{[]int{7, 1, 5, 3, 6, 4}},
			5,
		},
		{
			"2",
			args{[]int{7, 6, 5, 4, 3, 2, 1}},
			0,
		},
		{
			"3",
			args{[]int{1, 2, 3, 4, 5, 6, 7}},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
