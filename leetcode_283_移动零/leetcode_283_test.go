package leetcode_283_移动零

import (
	"testing"

	"gotest.tools/assert"
)

func TestMoveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{nums: []int{0, 1, 0, 3, 12}},
			want: []int{1, 3, 12, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			old := tt.args.nums
			moveZeroes(old)
			assert.DeepEqual(t, old, tt.want)
		})
	}
}
