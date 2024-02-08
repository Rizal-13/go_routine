package go_routine

import (
	"fmt"
	"testing"
	"time"
)

// Testing Channel
func Test_Channel(t *testing.T) {
	channel := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		channel <- "insun crypto 47"
	}()

	data := <-channel
	fmt.Println(data)
	close(channel)
}

func Test_Channel2(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(1 * time.Second)
		channel <- "insun crypto 1947"
		fmt.Println("Test Done")
	}()

	data := <-channel
	fmt.Println(data)
}

// Channel As Parameter
func GiveMeRespon(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "insun crypto"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeRespon(channel)

	data := <-channel
	fmt.Println(data)
}

// In Out Channel
func OnlyIn(channel chan<- string) {
	time.Sleep(1 * time.Second)
	channel <- "insun crypto 47"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}
