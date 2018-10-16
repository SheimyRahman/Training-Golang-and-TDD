package main

import "testing"

func TestFizzBuzz_3(t *testing.T) {
	got := fizzbuzz(3)
	want := "Fizz"

	if got != want {
		t.Errorf("fizzbuzz(3) \n got: %v \n want: \n%v", got, want)
	}
}
func TestFizzBuzz_5(t *testing.T) {
	got := fizzbuzz(5)
	want := "Buzz"

	if got != want {
		t.Errorf("fizzbuzz(5) \n got: %v \n want: \n%v", got, want)
	}
}

func TestFizzBuzz_4(t *testing.T) {
	got := fizzbuzz(4)
	want := "4"

	if got != want {
		t.Errorf("fizzbuzz(4) \n got: %v \n want: \n%v", got, want)
	}
}

func TestFizzBuzz_15(t *testing.T) {
	got := fizzbuzz(15)
	want := "FizzBuzz"

	if got != want {
		t.Errorf("fizzbuzz(15) \n got: %v \n want: \n%v", got, want)
	}
}

func TestFizzBuzzT_13(t *testing.T) {
	got := fizzbuzz(13)
	want := "13"

	if got != want {
		t.Errorf("fizzbuzz(13) \n got: %v\n want: %v\n", got, want)
	}
}

func TestFizzBuzz_30(t *testing.T) {
	got := fizzbuzz(30)
	want := "FizzBuzz"

	if got != want {
		t.Errorf("fizzbuzz(30) \n got: %v\n want: %v\n", got, want)
	}
}

func TestFizzBuzz_57(t *testing.T) {
	got := fizzbuzz(57)
	want := "Fizz"

	if got != want {
		t.Errorf("fizzbuzz(57) \n got: %v\n want: %v\n", got, want)
	}
}
