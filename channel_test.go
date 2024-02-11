package go_routine

import (
	"fmt"
	"strconv"
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

func TestOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "insun"
	channel <- "crypto"
	channel <- "master"

	// fmt.Println(<-channel)
	// fmt.Println(<-channel)
	// fmt.Println(<-channel)

	fmt.Println(cap(channel))
	fmt.Println(len(channel)) // harus tidak ditampilkan datanya agar dapat diketahui jumlah data di dalam buffer
	fmt.Println("Selesai")
}

func TestBufferedChannel1(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "insun"
		channel <- "crypto"
		channel <- "master"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
	fmt.Println(cap(channel))
	fmt.Println(len(channel))
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke" + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}
	fmt.Println("Done")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeRespon(channel1)
	go GiveMeRespon(channel2)

	counter := 0

	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}

}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeRespon(channel1)
	go GiveMeRespon(channel2)

	counter := 0

	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Waiting Data")
		}
		if counter == 2 {
			break
		}
	}

}
