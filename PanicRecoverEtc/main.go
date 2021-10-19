package main

import (
	"fmt"
	"log"
)

func main(){
	//Recover
	defer recovFuncy()

	//Defers//
	exampleDefer()

	variableDefer()

	fmt.Println()

	//Panics//
	issue1() //Note: GoLang will rarely stop a program when there is an error other than really bad problems, like dividing by zero

	panic("Or mannually panic like this") //Note: defer statements are ran before a panic statement
	
}

func issue1(){
	defer recovFuncy()

	c, b := 1, 0
	ans := c / b //cannot divide by zero, therefore panic
	fmt.Println(ans) //note that this will not run because the program panicked, but anything outside of this function will be recovered
}

func recovFuncy(){
	if err:= recover(); err != nil { //this defer will recover the program
		log.Println("Error: ", err)
	}
}

func exampleDefer(){
	//Defer means it will execute at the end of the function, before it returns
	defer fmt.Println(1) //defering multiple statements will execute them in a LIFO order
	defer fmt.Println(2) //defer statements take a function call, not a function itself
	fmt.Println(3)
}

func variableDefer(){
	a := "start"
	defer fmt.Println(a) //Note: defer takes the parameter value at the time of the defer
	a = "end"
}