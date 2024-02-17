package go_routine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGomaxProc(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()

	}

	totalcpu := runtime.NumCPU()
	fmt.Println("CPU", totalcpu)

	// runtime.GOMAXPROCS(15) (untuk merubah thread)
	totalthread := runtime.GOMAXPROCS(-1)
	fmt.Println("Thread", totalthread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Go Routine", totalGoroutine)

	group.Wait()
}
