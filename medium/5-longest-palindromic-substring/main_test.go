package main

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{s: "babad"},
			want: "bab",
		},
		{
			name: "Example 2",
			args: args{s: "cbbd"},
			want: "bb",
		},
		{
			name: "Example 3",
			args: args{s: "a"},
			want: "a",
		},
		{
			name: "Example 4",
			args: args{s: "ac"},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{s: "aba"},
			want: true,
		},
		{
			name: "Example 2",
			args: args{s: "bb"},
			want: true,
		},
		{
			name: "Example 3",
			args: args{s: "a"},
			want: true,
		},
		{
			name: "Example 4",
			args: args{s: "ac"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPalindrome(tt.args.s); got != tt.want {
				t.Errorf("checkPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
