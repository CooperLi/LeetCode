package leetcode_0035_搜索插入位置

import (
	"testing"
)

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TC1",
			args: args{
				nums:   []int{1,2,3},
				target: 0,
			},
			want: 0,
		},
		{
			name: "TC2",
			args: args{
				nums:   []int{1,3,5,6},
				target: 7,
			},
			want: 4,
		},
		{
			name: "TC3",
			args: args{
				nums:   []int{1,3,5,6},
				target: 3,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
