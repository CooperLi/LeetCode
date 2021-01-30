package leetcode_394_字符串解码

import (
	"testing"
)

func Test_decodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{"3[a]2[bc]"},
			want: "aaabcbc",
		},
		{
			name: "2",
			args: args{"3[a2[c]]"},
			want: "accaccacc",
		},
		{
			name: "3",
			args: args{"4[a3[b2[c2[d]]]]"},
			want: "abcddcddbcddcddbcddcddabcddcddbcddcddbcddcddabcddcddbcddcddbcddcddabcddcddbcddcddbcddcdd",
		},
		{
			name: "4",
			args: args{"1[a]"},
			want: "a",
		},
		{
			name: "5",
			args: args{"1[abc]1[abc]1[abc]1[abc]"},
			want: "abcabcabcabc",
		},
		{
			name: "6",
			args: args{"100[old]"},
			want: "leetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcode",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeString(tt.args.s); got != tt.want {
				t.Errorf("decodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
