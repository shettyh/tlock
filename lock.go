package tlock

import (
	"sync"
	"time"
)

type Lock interface {
	sync.Locker
	TryLock() bool
	TryLockWithTimeout(duration time.Duration) bool
}

type lock struct {
	lockChan chan struct{}
	locked bool
}

func New() Lock {
	// Create the channel with size 1

	return &lock{lockChan: make(chan struct{}, 1)}
}

func (l *lock) TryLock() bool {
	select {
	case l.lockChan <- struct{}{}:
		l.locked = true
		return true
	default:
		// Failed to Acquire l
		return false
	}
}

func (l *lock) TryLockWithTimeout(timeout time.Duration) bool {
	// fast path
	if l.TryLock() {
		return true
	}

	// slow path
	select {
	case l.lockChan <- struct{}{}:
		l.locked = true
		return true
	case <-time.After(timeout):
		if !l.locked && len(l.lockChan) == 1{
			l.locked = true
			return true
		}
		return false
	}


}

// lock is blocking call, waits for other lockChan to be released
func (l *lock) Lock() {
	l.lockChan <- struct{}{}
	l.locked = true
}

func (l *lock) Unlock() {
	select {
	case <-l.lockChan:
	default:
	}
	l.locked = false
	return
}
