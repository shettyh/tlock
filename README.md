# TLock
![Build](https://github.com/shettyh/tlock/workflows/Go/badge.svg?branch=master)

Simple Golang Lock with Timeout support.

## Install
- `go get github.com/shettyh/tlock`

## How to use

```go
// Create lock
tlock := tlock.New()

//blocking Lock/unlock
tlock.Lock()
defer tlock.Unlock()


// non-blocking lock/unlock
if tlock.TryLock() {
    defer tlock.Unlock()
    ...
}


// block lock/unlock with timeout
if tlock.TryLockWithTimeout(time.Seconds * 10 ) {
    defer tlock.Unlock()
    ...
}
```

For detailed example please check the examples folder

