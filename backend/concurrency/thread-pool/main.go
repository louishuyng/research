package main

import (
	"fmt"
	"sync"
	"time"
)

type Job func()

type Pool struct {
	workQueue chan Job
	wg        sync.WaitGroup
}

func (p *Pool) AddJob(job Job) {
	p.workQueue <- job
}

func (p *Pool) Wait() {
	close(p.workQueue)
	p.wg.Wait()
}

func NewPool(size int) *Pool {
	p := Pool{
		workQueue: make(chan Job, size),
	}
	p.wg.Add(size)
	for i := 0; i < size; i++ {
		go func() {
			defer p.wg.Done()

			for job := range p.workQueue {
				job()
			}
		}()
	}
	return &p
}

func main() {
	pool := NewPool(5)

	for i := 0; i < 30; i++ {
		fmt.Println(i)
		job := func() {
			time.Sleep(1 * time.Second)
			fmt.Printf("Job completed\n")
		}

		pool.AddJob(job)
	}

	pool.Wait()
}
