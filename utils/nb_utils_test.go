package utils

import "testing"

func BenchmarkConcat(b *testing.B) {
	examples := []string{"aa", "bb", "my name is nilay", "hello you", "go to heaven"}
	for i := 0; i < b.N; i++ {
		concat(examples)
	}
}
