package tlock

import (
	"time"
)

type Lock struct {
	lock chan struct{}
}

func New() *Lock {
	return &Lock{make(chan struct{}, 1)}
}

func (lock *Lock) TryLock() bool {
	select {
	case lock.lock <- struct{}{}:
		return true
	default:
		// Failed to Acquire lock
		return false
	}
}

func (lock *Lock) TryLockWithTimeout(timeout time.Duration) bool {
	if lock.TryLock() {
		return true
	}

	// Blocking send with timeout
	select {
	case lock.lock <- struct{}{}:
		return true
	case <-time.After(timeout):
		// Failed to Acquire lock
		return false
	}
}

// Blocking call
func (lock *Lock) Lock() {
	lock.lock <- struct{}{}
}

func (lock *Lock) Unlock() {
	select {
	case <-lock.lock:
		return
	default:
		return
	}
}
