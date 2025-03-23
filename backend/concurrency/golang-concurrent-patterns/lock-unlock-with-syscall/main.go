package lockunlockwithsyscall

import "sync/atomic"

type FutextLock int32

func (f *FutextLock) Unlock() {
	oldValue := atomic.SwapInt32((*int32)(f), 0)

	if oldValue == 2 {
		// futex_wake(addr, count) wakes up suspended executions (threads and processes)
		// that are waiting on the address specified.
		futext_wakeup((*int32)(f), 1)
	}
}

func (f *FutextLock) Lock() {
	if !atomic.CompareAndSwapInt32((*int32)(f), 0, 1) {
		for atomic.SwapInt32((*int32)(f), 2) != 0 {
			// futex_wait(addr, value), we specify a memory address and a value.
			// If the value at the memory address is equal to the specified parameter value,
			// the execution of the caller is suspended and placed at the back of a queue.
			futex_wait((*int32)(f), 2)
		}
	}
}
