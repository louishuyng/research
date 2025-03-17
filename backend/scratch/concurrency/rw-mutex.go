package main

import "sync"

type ReadWriteMutex struct {
	readersCounter int
	writerWaiting  int
	writerActive   bool
	cond           *sync.Cond
}

func NewReadWriteMutex() *ReadWriteMutex {
	return &ReadWriteMutex{
		cond: sync.NewCond(&sync.Mutex{}),
	}
}

func (m *ReadWriteMutex) RLock() {
	m.cond.L.Lock()

	for m.writerActive || m.writerWaiting > 0 {
		m.cond.Wait()
	}
	m.readersCounter++

	m.cond.L.Unlock()
}

func (m *ReadWriteMutex) RUnlock() {
	m.cond.L.Lock()

	m.readersCounter--
	if m.readersCounter == 0 {
		m.cond.Broadcast()
	}

	m.cond.L.Unlock()
}

func (m *ReadWriteMutex) Lock() {
	m.cond.L.Lock()

	m.writerWaiting++
	for m.writerActive || m.readersCounter > 0 {
		m.cond.Wait()
	}
	m.writerWaiting--
	m.writerActive = true

	m.cond.L.Unlock()
}

func (m *ReadWriteMutex) Unlock() {
	m.cond.L.Lock()

	m.writerActive = false
	m.cond.Broadcast()

	m.cond.L.Unlock()
}
