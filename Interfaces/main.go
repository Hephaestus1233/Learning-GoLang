package main

import (
	"fmt"
)

func main() {
	//var i interface{} = 0 //is an empty interface, and every type implements it even primitives

	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))
}

type Writer interface {
	Write([]byte) (int, error) //Just defined randomly
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type ConsoleWriter struct {}

func (cw ConsoleWriter) Write(data []byte) (int, error) { //implicitly links ConsoleWriter to Writer Interface
	n, err := fmt.Println(string(data))
	return n, err
}

//Note: the function Write for ConsoleWriter receives a value, but if it received a pointer/reference the object itself would need to be a pointer to implement it
