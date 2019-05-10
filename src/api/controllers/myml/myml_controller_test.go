package myml

import (
	"net/http"
	"testing"
)

func BenchmarkMain100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		http.Get("http://localhost:8080/myml/1")

	}
}
