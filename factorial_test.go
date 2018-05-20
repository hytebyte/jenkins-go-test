package main

import "testing"

func TestFactorialReturnsCorrectResults(t *testing.T) {
	cases := []struct{ n, result int }{
		{n: 0, result: 1},
		{n: 1, result: 1},
		{n: 2, result: 2},
		{n: 3, result: 6},
		{n: 4, result: 24},
		{n: 5, result: 120},
		{n: 6, result: 720},
	}

	for i, c := range cases {
		result := factorial(c.n)
		if result != c.result {
			t.Errorf("Case %d: Expected %d! == %d, got %d",
				i, c.n, c.result, result)
		}
	}
}

func TestFactorialPanicsOnInvalidInputs(t *testing.T) {
	if !panicked(func() { factorial(-1) }) {
		t.Error("Expected -1! to panic, but it didn't")
	}
}

func panicked(f func()) (panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()

	f()
	return
}
