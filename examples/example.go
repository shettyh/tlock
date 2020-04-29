package main

import (
	"fmt"
	"github.com/shettyh/tlock"
	"time"
)

type SingleAccessType struct {
	lock tlock.Lock
}

func New() *SingleAccessType {
	return &SingleAccessType{tlock.New()}
}

func (ac *SingleAccessType) TryAccess(timeout time.Duration) bool {
	if ac.lock.TryLockWithTimeout(time.Second * 5) {
		defer ac.lock.Unlock()
		fmt.Println("Access successful")
		return true
	}

	fmt.Println("Resource is busy not able access")
	return false
}

func main() {
	accessType := New()
	accessType.TryAccess(10 * time.Second)
}
