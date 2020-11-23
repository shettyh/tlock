package tlock

import (
	"sync"
	"sync/atomic"
	"time"
)

type Lock interface {
	sync.Locker
	TryLock() bool
	TryLockWithTimeout(duration time.Duration) bool
}

type lock struct {
	lockChan chan struct{}
	locked int32
}

func New() Lock {
	// Create the channel with size 1

	return &lock{lockChan: make(chan struct{}, 1)}
}

func (l *lock) TryLock() bool {
	select {
	case l.lockChan <- struct{}{}:
		atomic.StoreInt32(&l.locked, 1)
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
		atomic.StoreInt32(&l.locked, 1)
		return true
	case <-time.After(timeout):
		if atomic.LoadInt32(&l.locked) == 0 && len(l.lockChan) == 1{
			atomic.StoreInt32(&l.locked, 1)
			return true
		}
		return false
	}


}

// lock is blocking call, waits for other lockChan to be released
func (l *lock) Lock() {
	l.lockChan <- struct{}{}
	atomic.StoreInt32(&l.locked, 1)
}

func (l *lock) Unlock() {
	select {
	case <-l.lockChan:
	default:
	}
	atomic.StoreInt32(&l.locked, 0)
	return
}