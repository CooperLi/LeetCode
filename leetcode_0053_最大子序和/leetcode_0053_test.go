package leetcode_0053_最大子序和

import (
    "testing"
)

func Test_maxSubArray(t *testing.T) {
    type args struct {
        nums []int
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name:"1",
            args:args{[]int{-2,1,-3,4,-1,2,1,-5,4}},
            want:6,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := maxSubArray(tt.args.nums); got != tt.want {
                t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
            }
        })
    }
}

func Test_maxSubArray1(t *testing.T) {
    type args struct {
        nums []int
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name:"1",
            args:args{[]int{-2,1,-3,4,-1,2,1,-5,4}},
            want:6,
        },
        {
            name:"2",
            args:args{[]int{1}},
            want:1,
        },
        {
            name:"3",
            args:args{[]int{-2,1}},
            want:1,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := maxSubArray(tt.args.nums); got != tt.want {
                t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
            }
        })
    }
}
