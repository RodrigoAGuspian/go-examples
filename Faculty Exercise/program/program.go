package program

import (
	"fmt"
	"main/student"
)

type Program struct {
	Code     int
	Name     string
	Year     int
	Students []student.Student
}

const RECENT_YEAR = 2023
const NUM_PROG = 2

func createProgram(number int) Program {
	var program Program

	fmt.Printf("----Program #%d---- \n", number)
	for {
		fmt.Println("Input Program's code: ")
		fmt.Scanln(&program.Code)
		if program.Code <= 0 {
			fmt.Println("Error, input Program's code again!")
		} else {
			break
		}

	}

	fmt.Println("Input Program's name: ")
	fmt.Scanln(&program.Name)

	for {
		fmt.Println("Input Program's year of creation: ")
		fmt.Scanln(&program.Year)
		if program.Year <= 0 || program.Year > RECENT_YEAR {
			fmt.Println("Error, input Program's year of creation again!")
		} else {
			break
		}

	}

	program.Students = student.CreateAllStudents()

	//PrintProgram(program, number)

	return program

}

func CreateAllPrograms() []Program {
	var programs []Program

	for i := 0; i < NUM_PROG; i++ {
		programs = append(programs, createProgram(i))
	}

	return programs
}

func printProgram(program Program, number int) {

	fmt.Printf("----Program #%d---- \n", number)
	fmt.Printf("Code: %d; Name: %s; Year of creation: %d;\n", program.Code, program.Name, program.Year)

	student.PrintAllStudents(program.Students)
}

func PrintAllPrograms(program []Program) {
	for i := 0; i < NUM_PROG; i++ {
		printProgram(program[i], i)
	}
}
