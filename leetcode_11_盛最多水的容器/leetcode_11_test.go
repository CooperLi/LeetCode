package leetcode_11_盛最多水的容器

import (
    "testing"
)

func Test_maxArea(t *testing.T) {
    type args struct {
        height []int
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
          name:"1",
          args:args{[]int{1,8,6,2,5,4,8,3,7}},
          want:49,
        },
        {
            name:"2",
            args:args{[]int{1,8,6}},
            want:6,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := maxArea(tt.args.height); got != tt.want {
                t.Errorf("maxArea() = %v, want %v", got, tt.want)
            }
        })
    }
}