package go_routine

import (
	"fmt"
	"testing"
	"time"
)

func SayHello() {
	fmt.Println("Hello")
}

func TestSayHello(t *testing.T) {
	go SayHello()
	fmt.Println("Test Done")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoRoutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)

	}
	time.Sleep(5 * time.Second)
}

// func TestManyGoRoutine(t *testing.T) {
// 	for i := 0; i < 100000; i++ {
// 		DisplayNumber(i)
// 	}
// }
