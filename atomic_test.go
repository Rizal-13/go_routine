package go_routine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestAtomic(t *testing.T) {

	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		go func() {
			group.Add(1)
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println("Counter :", x)
}

func TestAtomic2(t *testing.T) {

	x := 0
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x = x + 1
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter :", x)
}
