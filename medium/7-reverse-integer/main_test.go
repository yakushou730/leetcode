package main

import "testing"

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
			name: "Example 1",
			args: args{x: 123},
			want: 321,
		},
		{
			name: "Example 2",
			args: args{x: -123},
			want: -321,
		},
		{
			name: "Example 3",
			args: args{x: 120},
			want: 21,
		},
		{
			name: "Example 4",
			args: args{x: 0},
			want: 0,
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
