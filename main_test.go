package main

import (
	"reflect"
	"testing"
)

func TestSieve(t *testing.T) {
	type args struct {
		max int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{max: 4},
			want: []int{2, 3},
		},
		{
			name: "test2",
			args: args{max: 5},
			want: []int{2, 3, 5},
		},
		{
			name: "test3",
			args: args{max: 13},
			want: []int{2, 3, 5, 7, 11, 13},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sieve(tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sieve() = %v, want %v", got, tt.want)
			}
		})
	}
}
