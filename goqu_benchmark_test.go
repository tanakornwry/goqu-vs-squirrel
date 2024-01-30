package main

import (
	"testing"
)

func TestGoqcSelectWithPGPlaceholder(t *testing.T) {
	//GoqcSelectWithPGPlaceholder()
	//GoquSelectSimple()
	//GoquInsertWithPGPlaceholder()
	//GoquUpdateWithPGPlaceholder()
	//GoquDeleteWithPGPlaceholder()
	GoquComplexSelectWithPGPlaceHolder()
}

func BenchmarkGoquSelectSimple(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		GoquSelectSimple()
	}
}
