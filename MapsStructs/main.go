package main

import (
	"fmt"
	//"reflect"
)

//Struct Example
type location struct{
	population int
	name string
}

type state struct{ //Capitalized first letters export the variable/struct (as well as internal variables if they to have their first letter capitalized)
	population int //`required max: "100"` //That's a tag to a value, accessible by using the reflect package to use var.FieldByName("population")
	name string 
	cities []string
	//etc, etc
}

type country struct{
	//state //This is how you do inheritance with structures, known as embedding
	states []state 
	name string
}

func main() {
	//Maps//
	statePopulations := make(map[string]int) //this is one way to make it
	statePopulations = map[string]int{ //another way to make (if you use :=) or to assign
		"California": 30000, //Otherwise known as dictionaries
		"Texas":      123123,
		"Suffering":  123123,
	}

	//Note: Cannot use arrays as keys
	statePopulations["Georgia"] = 123456
	delete(statePopulations, "Suffering")

	fmt.Println(statePopulations) //Note: Maps are not guaranteed to be in the same order as assignment on get
	fmt.Println(statePopulations["Georgia"])
	fmt.Printf("Amount of Populations: %v\n", len(statePopulations))

	value, success := statePopulations["Texs"] //Misspelled or non-existent entries return primitive default value

	fmt.Println(value, success) //However the success value will return false because it doesn't exist in the map

	//Note: Variable copying only copies the reference to the memory
	mapCopy := statePopulations
	mapCopy["Georgia"] = 12

	fmt.Println(statePopulations)
	fmt.Println(mapCopy)

	fmt.Println()

	//Structs//
	newYork := state{ //You do have the option to not include the names of each var in the struct and use positional declaration, but that is unmaintainable
		population : 123123,
		name: "New York",
		cities: []string{
			"New York",
		},
	}

	fmt.Println(newYork)
	fmt.Println(newYork.name, "has population:" ,newYork.population)

	fmt.Println()

	//Anonymous Structs
	position := struct{x int; y int}{x: 0, y: 0} //declare struct{declare variables}{initialize variables}

	fmt.Println(position)
	fmt.Printf("(%v, %v)\n", position.x, position.y)
	//Note: Structures are duplicated, not referenced, when assigned. A change to a copy will not affect the original
}