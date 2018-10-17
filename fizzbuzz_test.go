package main

import (
	"strconv"
	"testing"
)

var tests = []struct {
	number int
	want   string
}{
	{1, "1"},
	{3, "Fizz"},
	{4, "4"},
	{5, "Buzz"},
	{15, "FizzBuzz"},
	{30, "FizzBuzz"},
	{57, "Fizz"},
}

func TestFizzBuzz(t *testing.T) {
	for _, tc := range tests {
		t.Run(strconv.Itoa(tc.number), func(t *testing.T) {
			got := fizzbuzz(tc.number)

			if got != tc.want {
				t.Errorf("fizzbuzz(%d) \n got: %v \n want: \n%v", tc.number, got, tc.want)
			}
		})
	}
}
