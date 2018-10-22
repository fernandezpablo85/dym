package main

import (
	"testing"
)

func BenchmarkCorrections(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Corrections("juan fernand")
	}
}

func TestManuel(t *testing.T) {
	manuel := Corrections("mnuel")

	if manuel[0] != "manuel" {
		t.Errorf("should be 'manuel' but was %s\n", manuel[0])
	}
}
