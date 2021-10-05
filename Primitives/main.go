package main

import "fmt"

func main() {

	//Booleans//
	//var n bool = true

	n := 1 == 1
	var m bool //Every Variable in Go is auto initialized to a default value if not assigned, bool is initialized as false

	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", m, m)

	//Integers//
	//Any number type is default assigned 0
	var ( //Signed Integer Types
		standardSize int //Standard Integers are guaranteed a minimum of 32 bits, but can have more depending on OS
		size1        int8
		size2        int16
		size3        int32
		size4        int64
		example1     int = 10 //Binary is 1010
		example2     int = 3  //Binary is 0011
	)

	var ( //Unsigned Integer Types
		ustandardsize uint
		usize1        uint8
		//. . .
		//Note: GoLang does not have a uint64
	)
	fmt.Println(standardSize, size1, size2, size3, size4, ustandardsize, usize1)

	//Operations//
	//All normal operators apply to these numbers: (+ - * / %)
	/*Notes:
	1. Dividing two integers returns an integer and may lose information, type cast to float in order to get the decimal information
	2. Different int/number size types cannot be added together (will result in error) unless you do explicit type casting
	*/

	//Bit-wise operators
	fmt.Println(
		example1&example2,  //comparing which bits are 1 in both numbers, so therefore 0010 = 2
		example1|example2,  //comparing which bits are 1 in either number, so therefore 1011 = 3
		example1^example2,  //comparing which bits are 1 in one number but not the other, so therefore 1001 = 9
		example1&^example2, //comparing which bits are 0 in both numbers, so therefore 0100 = 8 (Operator can be considered the opposite of the bit-and)
		example1<<3,        //shifts the bits 3 to the left (multiplies number by 2^3)
		example1>>3,        //shifts the bits 3 to the right (divides number by 2^3)
	)

	//Floating-Point Numbers//
	var (
		rootbear float32
		orThis   float64
		//there does not exist a standard 'float' value
	)

	decimal1 := 3.  //Auto decides to use float64 rather than float32
	decimal1 = 3.14 //none of these will error because they are all considered floating point types
	decimal1 = 13.7e72
	decimal1 = 2.1e14 //Capitalizing E will just get auto-corrected to lowercase e by Go Linter

	fmt.Printf("%v, %T\n", rootbear, rootbear)
	fmt.Printf("%v, %T\n", orThis, orThis)
	fmt.Printf("%v, %T\n", decimal1, decimal1) //NOTE: Cannot use Modulo (%) operator on floats

	//Complex Numbers//
	var imaginary complex64 = 2i
	imaginary = 1 + 2i

	var anotherOne complex64 = complex(5, 12) //first is real, second is imaginary

	fmt.Printf("%v, %T\n", imaginary, imaginary)
	fmt.Println(imaginary + anotherOne) //All except modulo work on complexes

	//Extracting the components
	fmt.Printf("%v, %T\n", real(imaginary), real(imaginary)) //complex64 returns an int32
	fmt.Printf("%v, %T\n", imag(imaginary), imag(imaginary)) //complex128 returns an int64

	//Strings//
	aNormalString := "this is a string"                           //Strings are immutable
	aNormalString = aNormalString + " and this is another string" //you can add strings together

	asciiConversion := []byte(aNormalString)

	fmt.Printf("%v, %T\n", aNormalString, aNormalString)
	fmt.Printf("%v, %T\n", aNormalString[2], aNormalString[2])         //returns 105, uint5 for the ASCII table id
	fmt.Printf("%v, %T\n", string(aNormalString[2]), aNormalString[2]) //returns i, instead
	fmt.Printf("%v, %T\n", asciiConversion, asciiConversion)

	//Runes//
	//var r rune
	rExample := 'a' //Notice the assignment with single quotes

	fmt.Printf("%v, %T\n", rExample, rExample) //returns 97, int32
	//Note: Runes are really a type alias for int32s, basically int32 == rune
}
