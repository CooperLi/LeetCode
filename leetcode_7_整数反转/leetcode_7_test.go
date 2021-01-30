package leetcode_7_整数反转

import (
    "testing"
)

func Test_reverse(t *testing.T) {
    type args struct {
        x int
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name:"1",
            args:args{1534236469},
            want:0,
        },
        {
            name:"2",
            args:args{123},
            want:321,
        },
        {
            name:"3",
            args:args{-1230},
            want:-321,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := reverse(tt.args.x); got != tt.want {
                t.Errorf("reverse() = %v, want %v", got, tt.want)
            }
        })
    }
}