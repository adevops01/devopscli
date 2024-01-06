package main

import "test"

func Test_main(t *testing.T) {
	want := "Hello world"
	got := "Hellow world"

	if got != want {
		t.Errorf("%v got, but want %v", got, want)
	}
}