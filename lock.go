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
}

func New() Lock {
	// Create the channel with size 1
	return &lock{make(chan struct{}, 1)}
}

func (l *lock) TryLock() bool {
	select {
	case l.lockChan <- struct{}{}:
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

	endTime := time.Now().Add(timeout)
	// slow path
	for {
		select {
		case l.lockChan <- struct{}{}:
			return true
		default:
			// Failed to Acquire lock, check for timeout
			if endTime.Sub(time.Now()) <= 0.0 {
				return false
			}
		}
	}
}

// lock is blocking call, waits for other lockChan to be released
func (l *lock) Lock() {
	l.lockChan <- struct{}{}
}

func (l *lock) Unlock() {
	select {
	case <-l.lockChan:
		return
	default:
		return
	}
}
