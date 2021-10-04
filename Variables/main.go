package main

import (
	"fmt" //format (apparently)
	//"other/packages" can be imported here as well
)

//Global Variables
//var i int // Standard declaration

//Extra Important Note: lowercase variable names are not accessible to outside programs. The only way for programs that import your code to access variable is by UPPERCASE
var i int // Scoped to the package, not visible to external imports
var I int // Exported from the package and globally visible

//var i int is read as "Initializing a new variable named i which is an integer"

func main() {
	//Single-Variables//
	i = 42 // Standard Assignment

	var j int = 27 //Combined Declaration + Assignment
	var j2 float32 = 27

	k := 13 // Auto Declare Type + Assignment
	k2 := 13.
	//Note: Though GoLang can automatically assign the type of the variable at declaration, that variable type CANNOT change. We can't set n = "string" after.
	//Note: The Auto-Assign Variable operator ":=" cannot be used outside of a function for global variables like i

	fmt.Println(i, j, k) // Normal feed to output

	fmt.Printf("%v, %T\n", j, j) // Formatting output (Value, Type)

	fmt.Printf("%v, %T\n", j2, j2) // Printf doesn't include a newline at the end of the printing, so just include it

	fmt.Printf("%v, %T\n", k, k)

	fmt.Printf("%v, %T\n", k2, k2)

	//Multi-Line Variables//
	var (
		age        int  = 18 //BTW the Linter auto-corrects the formatting into... this, when you save in VSC. No I am not joking.
		isLearning bool = true
		// . . .
	)

	var name, friend string = "Henri Malahieude", "Chase Wheeler"

	fmt.Println("My name is", name, "and I am friends with", friend, ". I am ", age, " years old. Am I Learning?", isLearning)
	//Go Automatically puts spaces between the strings in the commas when using Println

	//Scope//
	var i int = 97 //i is already declared as a global variable, however we are shadowing it and redeclaring it as a local variable now
	//Note: You cannot redeclare this variable in the same scope (var i int = 20 on a new line) (i := 20 on a new line, ":=" is a declare and assign operator)
	fmt.Println("This variable is not the global variable! Local i is actually", i)

	//Type Casting//
	var randomDecimal float32 = float32(i) // If you do not explictly convert, you will get a compile time error
	randomDecimal = randomDecimal + 0.5
	fmt.Println("We have just cast the variable i into a float and added 0.5:", randomDecimal)

	var differentTypeAllTogether string = string(i) //The integer will be used to look up the character in the ASCII table

	fmt.Println("Type Casting i =", i, "into string:", differentTypeAllTogether) //importing "strconv" would allow you to use strcov.Itoa() function to get the actual string
}

//Other Note: The compiler WILL error if there are any unused variables.
