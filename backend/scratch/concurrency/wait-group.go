package main

import (
	"sync"
)

type WaitGroup struct {
	groupSize int
	cond      *sync.Cond
}

func NewWaitGroup() *WaitGroup {
	return &WaitGroup{
		cond: sync.NewCond(&sync.Mutex{}),
	}
}

func (wg *WaitGroup) Add(delta int) {
	wg.cond.L.Lock()
	wg.groupSize += delta
	wg.cond.L.Unlock()
}

func (wg *WaitGroup) Wait() {
	wg.cond.L.Lock()
	for wg.groupSize > 0 {
		wg.cond.Wait()
	}
	wg.cond.L.Unlock()
}

func (wg *WaitGroup) Done() {
	wg.cond.L.Lock()
	wg.groupSize--
	if wg.groupSize == 0 {
		wg.cond.Broadcast()
	}
	wg.cond.L.Unlock()
}
