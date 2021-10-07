package main

import (
	"fmt"
)

const namingConvention float32 = 21.1 //Shadowing applies

const ( //Enumerated Constants
	_ = iota //Note that underscoring as the name of a var/const will tell the compiler to throw that value away
	eC1 = iota + 3 //This moves the enumerator forward 3, without it the value would be 1,2,3 ... with it it is now 4,5,6 ...
	eC2 //iota is scoped to this const block, using it elsewhere would reset its value
	eC3 //Infers that we meant iota again
)

//Note that iota is most useful with bit-wise operations and storing data in bytes

func main() {
	//Regular Constants//
	const namingConvention = 36 //Please note that a constant cannot take a function as an input
	//Handles all primitives
	//Auto assigning constants doesn't take the normal ":=" operator, GoLang Compiler is lazy like that

	var interestingVariable int8 = 2 

	fmt.Printf("%v, %T\n", namingConvention, namingConvention) //This prints out 2, int

	//Note that this is interesting because constants can be added to different byte sizes of the same value, while variables will throw an error
	fmt.Printf("%v, %T\n", namingConvention + interestingVariable, namingConvention + interestingVariable) //this prints out 38, int8

	//Enumerated Constants//
	fmt.Printf("Enumerates: %v, %v, %v\n", eC1, eC2, eC3)
}
