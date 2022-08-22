package main

import "testing"

func TestGreet(t *testing.T) {
	t.Run("Greet", func(t *testing.T) {
		g := greet("World")
		if g != "Hi World" {
			t.Errorf("got %v, want Hi World", g)
		}
	})
}
