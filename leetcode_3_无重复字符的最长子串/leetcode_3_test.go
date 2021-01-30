package leetcode_3_无重复字符的最长子串

import (
    "testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
    type args struct {
        s string
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name: "1",
            args:args{s:"dvdf"},
            want:3,
        },
        {
            name: "2",
            args:args{s:"abc"},
            want: 3,
        },
        {
            name: "3",
            args:args{s:"pwwkew"},
            want: 3,
        },
        {
            name: "4",
            args:args{s:"au"},
            want: 2,
        },
        {
            name:"5",
            args:args{s:"abcabcbb"},
            want:3,
        },
        {
            name:"6",
            args:args{s:"tmmzuxt"},
            want:5,
        },
        {
            name:"7",
            args:args{s:"aabaab!bb"},
            want:3,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
                t.Errorf("LengthOfLongestSubstring() = %v, want %v", got, tt.want)
            }
        })
    }
}

