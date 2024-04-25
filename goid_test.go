package golocal

import (
	"testing"
	"time"
)

func TestGoId(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func() {
			t.Log(GoId())
		}()
	}
	time.Sleep(time.Second)
}

func BenchmarkGoId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoId()
	}
}

func BenchmarkXxx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(i)
	}
}
