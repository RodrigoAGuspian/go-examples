package faculty

import (
	"fmt"
	"main/program"
)

type Faculty struct {
	Code     int
	Name     string
	Programs []program.Program
}

func CreateFaculty() Faculty {
	var faculty Faculty
	fmt.Println("-----Faculty-----")
	for {
		fmt.Println("Input Faculty's code: ")
		fmt.Scanln(&faculty.Code)
		if faculty.Code <= 0 {
			fmt.Println("Error, input Faculty's code again!")
		} else {
			break
		}
	}
	fmt.Println("Input Faculty's name: ")
	fmt.Scanln(&faculty.Name)

	faculty.Programs = program.CreateAllPrograms()

	return faculty

}

func PrintFaculty(faculty Faculty) {
	fmt.Println("-----Faculty-----")
	fmt.Printf("Code: %d; Name: %s;\n", faculty.Code, faculty.Name)
	program.PrintAllPrograms(faculty.Programs)
	program.CountMaleStudentsInProgram(faculty.Programs)
}
