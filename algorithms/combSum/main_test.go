package main

import "testing"

func BenchmarkSteps(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stepPerms(35)
	}
}
