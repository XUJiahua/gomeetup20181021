package embintf

import (
	"testing"
)

// BenchmarkCachedCatRepo_Get-8   	100000000	        12.2 ns/op
func BenchmarkCachedCatRepo_Get(b *testing.B) {
	mCatRepo := NewDBCatRepo("127.0.0.1:3306")
	rCatRepo := NewCachedCatRepo(mCatRepo)

	benchmark(b, rCatRepo)
}

// BenchmarkDbCatRepo_Get-8   	    5000	    267805 ns/op
func BenchmarkDbCatRepo_Get(b *testing.B) {
	mCatRepo := NewDBCatRepo("127.0.0.1:3306")

	benchmark(b, mCatRepo)
}

func benchmark(b *testing.B, repo CatRepo) {
	for i := 0; i < b.N; i++ {
		repo.Get(1)
	}
}
