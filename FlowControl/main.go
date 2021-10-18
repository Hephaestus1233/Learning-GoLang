package main

import "fmt"

func main(){
	//If-Statements//
	if true {
		fmt.Println()
	}

	arbitraryMap := map[string]int{ "Rando": 123 }

	if rrr, ok := arbitraryMap["Rando"]; ok {
		fmt.Println(rrr) //Note rrr is only defined in this if-scope
	}

	var guess int = 12;
	if guess == 123 { //All other comparitors work, >, <, <=, >=, ==, !=
		fmt.Println("bruh")
	}

	if (guess == 123 || guess > 123){ //Or = ||, And = &&
		fmt.Println("broh!!!")
		fmt.Println(!true) //Not Operator is !
	} else if guess <= 123 {
		fmt.Println("this is an example of an else-if statement")
	} else {
		fmt.Println("this is an example of an else statement")
	}

	//Note: Short-Circuiting if-statements exists in GoLang
	//Note: When comparing for equality between floating point values, be aware that floats are estimations and not exact values

	fmt.Println()

	//Switch-Statements//

	//Standard
	switch 2{ //switch <tag> {
	case 1: //this value (1) will be compared to the <tag>
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("neither")
	}

	//Multi-Comparison
	switch 2{
	case 1, 5, 10:
		fmt.Println("1, 5, or 10")
	case 2, 4, 6: //Comparitors must be unique, cannot have 5, 1, or 10 in this "case"
		fmt.Println("2, 4, or 6")
	default:
		fmt.Println("homie, how")
	}

	//Initializor Example
	switch i := 2 + 3; i{ //initialize something in here, and use it as the tag
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("neither")
	}

	//If-Statement like
	i := 1
	switch {
	case i<=10:
		fmt.Println(10)
		fallthrough //this is an actual keyword, it will run the next case no matter if it evaluates the next check, if this case is true
	case i >= 20: //this allows for overlapping comparitors, compared to the multi-comparison switch statement
		fmt.Println(20)
	default:
		fmt.Println("Suffering from success")
	}

	//Type Switch
	var arb interface{} = 1

	switch arb.(type){
	case int:
		fmt.Println("int")
		break
		fmt.Println("This will not be evaluated")
	case float64:
		fmt.Println("float64")
	default:
		fmt.Println("Simply compares the type of the variable")
	}

	//Note: break keyword is implied in the switch blocks, as it is redundant in other languages
	//Note: for arrays to evaluate as equivalent, they must be both have the same type and *size*
}