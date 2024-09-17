package main

import (
	"fmt"
)

func main() {
	student := []string{}
	students := [][]string{}
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
	student = append(student, "David", "History", "B+")
	students = append(students, student)
	for i := 0; i < len(students); i++ {
		fmt.Println(students[i][0])
	}
}
