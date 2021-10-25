package main

import "fmt"

//Every Go program must have an entry point of a main function that returns no values
//Just like with variables, an uppercase First letter determines visibility outside the file
func main(){ //Dude, thank god, Go requires the open brace to be on the same line as the parantheses
	for i := 0; i < 5; i++{
		sayMessage("Hello Go!", i)
	}

	sayGreeting("Hello", "Person")

	s := sum(1, 2, 3, 4, 5, 6) //or as little and as big as you want
	fmt.Println("The sum is", s)

	s = sum1(1, 2, 3, 4, 5, 6)
	fmt.Println("This sum is", s)

	//We know this is going to error, and this assigns multiple variables. Note that since I don't use the answer, I mark it as unused to the compiler with an underscore
	_, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
	}

	//Anonymous function, runs once and can't be called otherwise
	func() {
		fmt.Println("Anon be like!")
	}() //required to have the call/invoke operator/syntax here

	for i := 0; i < 5; i++ {
		func(y int){ //Good practice to include a parameter because parallel threads may cause strange behavior
			fmt.Println(y)
		}(i) //not the inclusion of i into the invoke parameters
	}

	//Functions are types of variables as well
	f := func(){
		fmt.Println("Hello Go! 1")
	}

	f()

	var f2 func() = func() {
		fmt.Println("Hello Go! 2")
	}

	f2()

	var division func(float64, float64) (float64, error) //note that this function can only be accessed after it's initial declaration unlike the global function
	division = func(f1, f2 float64) (float64, error) {
		if (f2 == 0.0){
			return 0.0, fmt.Errorf("Cannot divide by zero")
		}
	
		return f1 / f2, nil
	}

	_, err2 := division(5.0, 0.0)
	if err2 != nil {
		fmt.Println(err2)
	}

	methodExample := greeter{
		greeting : "Hello",
		name : "Brother",
	}

	methodExample.greet() //Implicityly pass the greeter struct to the function
}

func sayMessage(msg string, idx int){
	fmt.Println(msg, idx)
}

//Syntactical Sugar
func sayGreeting(greeting, name string){ //no need to have string after each of them
	fmt.Println(greeting, name)
}

/*Note: When passing data into a function, it is passed by value and therefore copied into the function. 
		Editting the parameter value within the function won't change it outside unless it's a map or a slice (those are pointers)*/
//Note: Ofc, if you pass in pointers as parameters then you can change the value of the variable outside the function

func sum(values ... int) int{ //Variadic parameters (adding them all up into an array), must be the last parameter
	result := 0
	for _, v := range values {
		result += v;
	}

	return result
}
// (function keyword) (name) (parameters) (return type if needed) {
// func               sum   (values...int)  int                   {

//Possible return types: primitives, structs, even pointers to variables

//Syntactical Sugar
func sum1(values ... int) (result int){
	for _, v := range values {
		result += v;
	}
	return 
}

//Multiple return values
func divide(a, b float32) (float32, error){
	if (b == 0.0){
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}

	return a / b, nil
}

//Methods section//
type greeter struct{
	greeting string
	name string
}

func (g greeter) greet(){
	fmt.Println(g.greeting, g.name)
}