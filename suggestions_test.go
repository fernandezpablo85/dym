package main

import (
	"testing"
)

func BenchmarkSuggestions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Suggestion("juan fernand") // about 50µs
	}
}

var suggestions = []struct {
	orig     string
	expected string
}{
	{"andrea  fabiana", "andrea fabiana"},
	{"gisella alejandra", "gisella alejandra"},
	{"fantino", "santino"},
}

func TestSuggestions(t *testing.T) {
	for _, example := range suggestions {
		s := Suggestion(example.orig)
		if s != example.expected {
			t.Errorf("expected '%s' but was '%s'", example.expected, s)
		}
	}
}
