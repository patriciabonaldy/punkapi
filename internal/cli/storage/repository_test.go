package cli

import "testing"

func BenchmarkGetBeers(b *testing.B) {
	repo := NewRepository()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		repo.GetBeers()
	}
}
