package leetcode_435_无重叠区间

import (
	"testing"
)

func TestEraseOverlapIntervals(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}},
			want: 1,
		},
		{
			name: "2",
			args: args{[][]int{{1, 2}, {2, 6}, {1, 4}, {3, 7}, {4, 5}}},
			want:
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eraseOverlapIntervals(tt.args.intervals); got != tt.want {
				t.Errorf("eraseOverlapIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
