package deadlock

import (
	"fmt"
	"sync"
	"time"
)

func red(lock1, lock2 *sync.Mutex) {
	for {
		fmt.Println("Acquiring lock1")
		lock1.Lock()
		fmt.Println("Acquiring lock2")
		lock2.Lock()
		fmt.Println("Locks acquired")
		lock2.Unlock()
		lock1.Unlock()
		fmt.Println("Both Locks released")
	}
}

func blue(lock1, lock2 *sync.Mutex) {
	for {
		fmt.Println("Acquiring lock2")
		lock2.Lock()
		fmt.Println("Acquiring lock1")
		lock1.Lock()
		fmt.Println("Locks acquired")
		lock1.Unlock()
		lock2.Unlock()
		fmt.Println("Both Locks released")
	}
}

func main() {
	lock1 := &sync.Mutex{}
	lock2 := &sync.Mutex{}

	go red(lock1, lock2)
	go blue(lock1, lock2)

	time.Sleep(20 * time.Second)

	fmt.Println("Done")
}
