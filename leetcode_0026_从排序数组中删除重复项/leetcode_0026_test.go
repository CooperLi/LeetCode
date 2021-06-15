package leetcode_0026_从排序数组中删除重复项

import (
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{[]int{1, 2, 3, 3, 4}},
			4,
		},
		{
			"2",
			args{[]int{1, 2, 3, 4}},
			4,
		},
		{
			"3",
			args{[]int{1}},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
