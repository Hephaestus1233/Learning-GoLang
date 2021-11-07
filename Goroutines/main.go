package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
//using sync.Mutex and sync.RWMutex can lock information to prevent routines from interacting with it until another routine is done with it

func main(){
	runtime.GOMAXPROCS(6) //sets the maximum amount of threads to run. Default is set to the amount of cores on your cpu
	//Note: More threads per core can actually increase performance, just don't overwork the go scheduler
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1)) //output the max amount of threads we can use
	//Simple Concurrency//
	
	//Method 1
	go sayHello() //tells compiler to spin off a new "green" thread
	

	//Method 2
	var msg string = "Hello 2"
	go func(){
		fmt.Println(msg)
	}()
	msg = "actually it's time to say goodbye because of dependency issues/race condition"

	//Method 3
	var msg2 string = "Better Hello"
	go func(in string){
		fmt.Println(in)
	}(msg2)

	//since the functions above don't actually have enough time to run before the main function is "deleted" we can sleep the program
	time.Sleep(100 * time.Millisecond) //Note: This is absolutely terrible because it ties the go routines to the clock cycle. Instead let's use waitgroups

	//Adding shit to wait groups//
	wg.Add(1) //add 1, goroutine to wait on
	go func(){
		fmt.Println("1")
		wg.Done() //decrements the amount of routines that the waitgroup is waiting on
	}()

	wg.Add(1)
	go func(){
		fmt.Println("2")
		wg.Done()
	}()

	wg.Wait()
}

func sayHello(){
	fmt.Println("Hello!")
}

//If you want to detect race conditions in the code, use "go run -race [file location]", the -race flag will have the compiler see if it can find any race conditions