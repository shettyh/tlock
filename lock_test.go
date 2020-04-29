package tlock

import (
	"testing"
	"time"
)

func BenchmarkLock_Unlock(b *testing.B) {
	lock := New()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		lock.Lock()
		lock.Unlock()
	}
}

func BenchmarkLock_Unlock_Parallel(b *testing.B) {
	lock := New()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lock.Lock()
			lock.Unlock()
		}
	})
}

func BenchmarkTryLock_Unlock(b *testing.B) {
	lock := New()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		val := lock.TryLock()
		if val {
			lock.Unlock()
		}
	}
}

func BenchmarkTryLock_Unlock_Timeout(b *testing.B) {
	lock := New()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		val := lock.TryLockWithTimeout(time.Nanosecond * 2)
		if val {
			lock.Unlock()
		}
	}
}

func BenchmarkTryLock_UnlockParallel(b *testing.B) {
	lock := New()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			val := lock.TryLock()
			if val {
				lock.Unlock()
			}
		}
	})
}

func BenchmarkTryLock_UnlockParallel_Timeout(b *testing.B) {
	lock := New()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			val := lock.TryLockWithTimeout(time.Nanosecond * 2)
			if val {
				lock.Unlock()
			}
		}
	})
}
