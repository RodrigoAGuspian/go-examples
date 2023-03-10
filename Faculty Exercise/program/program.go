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

func CreateProgram(number int) {
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

	PrintProgram(program, number)

}

func PrintProgram(program Program, number int) {

	fmt.Printf("----Program #%d---- \n", number)
	fmt.Printf("Code: %d; Name: %s; Year of creation: %d;\n",
		program.Code, program.Name, program.Year)

	student.PrintAllStudents(program.Students)
}
