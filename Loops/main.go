package main

import "fmt"

func main(){
	//For-Loops//
	for i := 0; i < 5; i++ {
		fmt.Println(i);
		if i == 0 {
			i++; //Note: i is variable in this scope like any other and can be messed with
		}
	}

	for i, j := 0, 0; i < 5; i, j = i+1, j+2{ //multi-increment
		fmt.Println(i, j);
	}

	y := 0
	for ; y < 5; y++{ //Note: Cannot remov ethat first semi-colon even though it looks useless
		fmt.Println(y)
	}

	for i := 0; i < 3; {
		fmt.Println(i)
		i++; //Note: despite not needing to give the loop an incrementor, it will need to be incremented manually or it will loop indefinitely
	}

	x := 0
	for x < 5 { //Otherwise known as "syntactical sugar"
		fmt.Println(x);
		x++;
	}

	x = 0
	for { //Note: Infinite for loop will not be allowed to run unless there is a break keyword within
		fmt.Println(x);
		x++
		if x >= 5 {
			break
		}
	}

	for i := 0; i < 3; i++{
		if i == 2 {
			continue //Yup, it's just chuck testa
		}
		fmt.Println(i)
	}

Loop: //This is a label
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			fmt.Println(i * j)
			if i * j >= 3 {
				break Loop //Notifies the compiler that it wants to break all the way out to the label
			}
		}
	}

	s := []int{1, 2, 3}
	for k, v := range s{ //works for slices, arrays, maps, and even strings (though value will be a rune/int32)
		fmt.Println(k, v)
	}

	for _, v := range s{
		fmt.Println(v) //This is how you just get the value variable without getting an error from the compiler about an unused variable
	}
}