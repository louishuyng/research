package main

import "sync"

type Semaphore struct {
	permits int
	cond    *sync.Cond
}

func NewSemaphore(permits int) *Semaphore {
	return &Semaphore{
		permits: permits,
		cond:    sync.NewCond(&sync.Mutex{}),
	}
}

func (s *Semaphore) Acquire() {
	s.cond.L.Lock()

	for s.permits <= 0 {
		s.cond.Wait()
	}
	s.permits--

	s.cond.L.Unlock()
}

func (s *Semaphore) Release() {
	s.cond.L.Lock()

	s.permits++
	s.cond.Broadcast()

	s.cond.L.Unlock()
}
