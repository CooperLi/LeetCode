package leetcode_0046_全排列

import (
	"reflect"
	"testing"
)

func Test_permute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{[]int{1, 2, 3}},
			want: [][]int{
				[]int{1, 2, 3},
				[]int{1, 3, 2},
				[]int{2, 1, 3},
				[]int{2, 3, 1},
				[]int{3, 1, 2},
				[]int{3, 2, 1},
			},
		},
		// {
		// 	name: "2",
		// 	args: args{[]int{0, 1}},
		// 	want: [][]int{
		// 		[]int{0, 1},
		// 		[]int{1, 0},
		// 	},
		// },
		// {
		// 	name: "3",
		// 	args: args{[]int{5, 4, 6, 2}},
		// 	want: [][]int{
		// 		[]int{5, 4, 6, 2},
		// 		[]int{5, 4, 2, 6},
		// 		[]int{5, 6, 4, 2},
		// 		[]int{5, 6, 2, 4},
		// 		[]int{5, 2, 4, 6},
		// 		[]int{5, 2, 6, 4},
		// 		[]int{4, 5, 6, 2},
		// 		[]int{4, 5, 2, 6},
		// 		[]int{4, 6, 5, 2},
		// 		[]int{4, 6, 2, 5},
		// 		[]int{4, 2, 5, 6},
		// 		[]int{4, 2, 6, 5},
		// 		[]int{6, 5, 4, 2},
		// 		[]int{6, 5, 2, 4},
		// 		[]int{6, 4, 5, 2},
		// 		[]int{6, 4, 2, 5},
		// 		[]int{6, 2, 5, 4},
		// 		[]int{6, 2, 4, 5},
		// 		[]int{2, 5, 4, 6},
		// 		[]int{2, 5, 6, 4},
		// 		[]int{2, 4, 5, 6},
		// 		[]int{2, 4, 6, 5},
		// 		[]int{2, 6, 5, 4},
		// 		[]int{2, 6, 4, 5},
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permute(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
