package main

import (
	"testing"
)

func BenchmarkSQLBuilder(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		GoquBuildSQLSuit()
		//SQBuildSQLSuit()
	}
}
