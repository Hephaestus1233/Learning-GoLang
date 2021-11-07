package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	//The only way to make a channel
	ch := make(chan int) // ch is a channel that passes through integers
	
	wg.Add(2)
	go func(){
		i := <- ch //waiting on information from this channel
		fmt.Println(i)
		wg.Done()
	}()
	
	//Note: If there are no threads receiving data from the channel, sending threads will fall asleep and cause an error (deadlock)
	go func(){
		ch <- 42 //sending information through the channel, this sends a copy, this yields until another thread is yielded to receive
		wg.Done()
	}()
	wg.Wait()

	
	//Forcing the channel to be read only or write only//
	wg.Add(2)
	go func(ch1 <- chan int){ //can ONLY read information
		i := <- ch1
		fmt.Println(i)
		wg.Done()
	}(ch)
	
	go func(ch1 chan <- int){ //can ONLY send information
		ch1 <- 42
		wg.Done()
	}(ch)
	wg.Wait()

	//Buffered Channels//
	bch := make(chan int, 50) //second parameter is a buffer size
	wg.Add(2)
	go func(ch1 <- chan int){ //can ONLY read information
		i := <- ch1
		l, ok := <- ch1
		fmt.Println(i, l, ok) //manually check, the ok boolean will return whether or not the channel is open


		/*for {
			j, ok := <- ch1
			fmt.Println(ok)
			if ok{
				fmt.Println(j)
			} else {
				break
			}
		}*/ //There seems to be a deadlock issue with this, I should be able to use "range ch" and other loop methods to detect it being closed. I'll leave this like this until I understand this problem better

		wg.Done()
	}(bch)
	
	go func(ch1 chan <- int){ //can ONLY send information
		ch1 <- 42
		ch1 <- 12 //this won't error the program because the channel has a buffer of messages so it will hold that in its buffer until a thread asks for a value
		ch1 <- 21
		ch1 <- 51
		ch1 <- 123
		close(ch) //Note: This is so that the for range loop will finally end, without it there will be a deadlock
		//Note: once a channel is closed, it is closed forever and will not be able to be reopened. There won't even be a way to detect if the channel is closed on the sending side
		wg.Done()
	}(bch)
	wg.Wait()

}

//Note: "var channel = make(chan struct{})" has 0 memory size due to the empty struct, and is very efficient if you just want to send a single message/event