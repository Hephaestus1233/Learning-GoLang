package main

import "fmt"

func main(){


	//Normal Declarations//
	//Arrays are contiguous in memory, as defined by compiler (therefore more efficient than multiple variables)
	grades := [3]int{1, 2, 3} //name := [size]type{initialization values}
	gradesAutoSize := [...]int{1, 2, 3}
	var emptyGrades [3]int
	

	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Grades Auto: %v\n", gradesAutoSize)
	fmt.Printf("Grades Empty: %v\n", emptyGrades)
	fmt.Printf("Grades #1: %v\n", grades[0]) //standard access operation
	emptyGrades[0] = 1 //standard assignment operation

	fmt.Printf("Amount of Grades: %v, %T\n", len(grades), len(grades)) //built-in len function can be used to find the size of an array

	fmt.Println()

	//Multi-Dimensional Arrays//
	var identityMatrix [3][3]int = [3][3]int{ 
		{1, 0, 0}, 
		[3]int{0, 1, 0}, //though this notation is redudant
	}

	identityMatrix[2] = [3]int{0, 0, 1} //you do need the redundant notation here though

	fmt.Println(identityMatrix)

	//Note: Arrays are considered actual values, instead of simply pointers to other values
	gradeCopy := grades //So this will actually be completely copied to the gradeCopy value, it won't reference the location of the first as C++ may do
	grades[1] = 5
	fmt.Printf("Grade: %v\n", grades)
	fmt.Printf("Grade Copy: %v\n", gradeCopy)

	fmt.Println()

	//Slices//
	students := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"} //only difference with an array is that you don't initialize a size

	fmt.Printf("Amount of Students: %v, %T\n", len(students), len(students))
	fmt.Printf("Capacity of Students: %v, %T\n", cap(students), cap(students)) //Only works for slices

	//Note: Slices are simply references to memory locations, so setting another variable equal to another won't copy it over
	studentCopy := students //the opposite of arrays
	studentCopy[1] = "Z" //In arrays this would only apply to the copy, but with slices it applies to "both" because the copy isn't a copy but a memory pointer

	fmt.Printf("Students: %v\n", students)
	fmt.Printf("StudentCopy: %v\n", studentCopy)

	fmt.Println()

	//Slice Notation//
	literalCopy := students[:] //slice of all elements
	lastStudents := students[4:] //slice from 4th elements to end (copy starts from the 3rd index)
	firstStudents := students[:6] //slice first 6 elements (copy ends at index 5)
	middleStudents := students[3:6] //slice 4th, 5th, 6th elements
	//bracket + semicolon notation is [Inclusive:Exclusive]
	//Note: These are all pointers, a change in one will affect ALL others
	//Note: This notation can be used to move [Array -> Slice] and [Slice -> Slice]

	fmt.Println(literalCopy)
	fmt.Println(lastStudents)
	fmt.Println(firstStudents)
	fmt.Println(middleStudents)

	fmt.Println()

	//Make Function//
	bruh := make([]int, 3, 100) //creates a slice with a size of 3, and capacity of 100
	fmt.Println(bruh)
	fmt.Printf("Length of Bruh: %v\n", len(bruh))
	fmt.Printf("Capacity of Bruh: %v\n", cap(bruh))

	//Changing Size of Slice//
	bruh = append(bruh, 4, 5, 6) //I believe this only applies to slices, and adds 4, 5, 6 to the slice
	fmt.Printf("Length of Bruh: %v\n", len(bruh))
	fmt.Printf("Capacity of Bruh: %v\n", cap(bruh))
	//Note: Appending to a slice that has reached max size of capacity (5 size / 5 capacity) will start increasing the capacity by much more than intended and waste memory

	fmt.Println()

	//Concatenating Two Slices"//
	doubleSize := append(students, literalCopy...) //the ... operator will take each element of the slice and place them into the argument list of the function
	fmt.Println(doubleSize)

	fmt.Println()

	//Pop Operations//
	popFirst := students[1:] //Removes first element
	fmt.Println(popFirst) 
	popLast := students[:len(students)-1]
	fmt.Println(popLast)

	fmt.Println();

	//Removal Operations//
	removeMid := append(students[:3], students[4:]...) //Be careful with this though, there may be unexpected behavior if you still have references to the original slice
	fmt.Println(removeMid)
	fmt.Println(popFirst) //this ends with j, j. Even though it should be i, j because we removed one element in the middle
	fmt.Println(popLast)
}