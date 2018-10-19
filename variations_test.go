package main

import (
	"testing"
)

func TestVariations(t *testing.T) {
	vars := Variations("pablo")

	mustContain := func(str string) {
		if !contains(vars, str) {
			t.Errorf("variations must contain '%s'", str)
		}
	}

	mustContain("pblo")
	mustContain("peblo")
	mustContain("pabli")
	mustContain("pabloa")
	mustContain("pablo")
}

func contains(els []string, item string) bool {
	for _, el := range els {
		if el == item {
			return true
		}
	}
	return false
}

var result []string

func benchmarkVariations(s string, b *testing.B) {
	var vs []string
	for i := 0; i < b.N; i++ {
		vs = Variations(s)
	}
	result = vs
}

func BenchmarkVariationsFive(b *testing.B) {
	benchmarkVariations("pablo", b)
}

func BenchmarkVariationsTen(b *testing.B) {
	benchmarkVariations("Juan Carlo", b)
}

func BenchmarkVariationsTwentyOne(b *testing.B) {
	benchmarkVariations("marianela guillermina", b)
}

func BenchmarkVariationsThirtyFive(b *testing.B) {
	benchmarkVariations("san fernando del valle de catamarca", b)
}
