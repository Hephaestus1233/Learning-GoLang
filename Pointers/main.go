package main

import "fmt"

func main(){

	//Standard Variable Copying//
	var a int = 42
	b := a //Duplicated into the variable
	fmt.Println(a, b)
	a = 27 //Change here does not affect the variable b
	fmt.Println(a, b)
	fmt.Println()

	//Using Pointers (that point to places in memory)//
	var aa int = 42
	var bb *int = &aa //Stores the memory location that stores the contents of the variable aa

	fmt.Println(&aa, bb) //These are the same values (0xasdklalsdjl123, etc)
	fmt.Println(aa, *bb) //Ditto

	aa = 27 //This does change bb because bb is "pointing" to its location in memory
	fmt.Println(aa, *bb) //This is because they are both pointing to the same place in memory, just in a different way

	*bb = 12 //This does the same as above
	fmt.Println(aa, *bb)
	fmt.Println()

	//Ways to create a Pointer//
	i := 1
	var j *int = &i
	k := &a
	l := new(int) //or a different type, will hold a nil value though

	fmt.Println(i, j, k, l)

	//Note: the dereference operator (*) returns the value of a pointer/reference, the reference operator (&) returns the memory address of the variable
	//Note: the dereference operator (*) has a lower precedence than the dot (.) operator, meaning you'll need to put parenthesis around an object pointer you've dereferenced and want to access member functions on
	//Note: Or you can just do objectPointer.memberFunctionOrVar = etc, the compiler will interpret it for you

	//Note: cannot do any pointer arithmetic with the memory address you get for references/pointers
	//Note: If you desperately want pointer arithmetic, then import the "unsafe" package that way the compiler will allow you to mess with memory addresses in the code
	//Note: Slices and Maps are not copied to new variables but by pointer, so editting a var that is initialized from a slice or a map will change the original slice/map
}