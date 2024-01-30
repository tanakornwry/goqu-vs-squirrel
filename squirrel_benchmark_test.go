package main

import (
	"testing"
)

func TestSQSelectWithPgPlaceholder(t *testing.T) {
	//SQSelectWithPgPlaceholder()
	//SQSelectSimple()
	//SQInsertWithPGPlaceholder()
	//SQUpdateWithPGPlaceholder()
	//SQDeleteWithPGPlaceholder()
	SQComplexSelectWithPGPlaceHolder()
}

func BenchmarkSquirrelSelectSimple(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SQSelectSimple()
	}
}
