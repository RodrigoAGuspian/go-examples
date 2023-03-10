package student

import "fmt"

const NUM_STU = 3

type Student struct {
	Code     int
	Name     string
	LastName string
	Gender   rune
}

func CreateStudent(number int) Student {
	var student Student

	for {
		fmt.Println("Input Student's code: ")
		fmt.Scanln(&student.Code)
		if student.Code <= 0 {
			fmt.Println("Error, input Student's code again!")
		} else {
			break
		}

	}
	fmt.Printf("----Student #%d---- \n", number)
	fmt.Println("Input Student's name: ")
	fmt.Scanln(&student.Name)
	fmt.Println("Input Student's last name: ")
	fmt.Scanln(&student.LastName)
	for {
		fmt.Println("Input Student's gender (m or f): ")
		fmt.Scanf("%c\n", &student.Gender)

		if student.Gender != 'm' && student.Gender != 'f' {
			fmt.Println("Error, input Student's gender again!")
		} else {
			break
		}
	}

	return student
}

func CreateAllStudents() []Student {
	var students []Student

	for i := 0; i < NUM_STU; i++ {
		students = append(students, CreateStudent(i))
	}

	return students

}

func PrintStudent(student Student) {
	fmt.Printf("Code: %d; Name: %s; Last Name: %s; Gender: %c;\n",
		student.Code, student.Name, student.LastName, student.Gender)

}

func PrintAllStudents(students []Student) {
	for i := 0; i < NUM_STU; i++ {
		PrintStudent(students[i])
	}
}
