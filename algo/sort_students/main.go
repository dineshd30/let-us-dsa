package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Student struct {
	name         string
	mathGrade    int
	scienceGrade int
}

func (s *Student) getTotalGrade() int {
	return s.mathGrade + s.scienceGrade
}

func getStudentFromStr(str string) (*Student, error) {
	items := strings.Split(str, " ")
	mathGrade, err := strconv.Atoi(items[1])
	if err != nil {
		fmt.Println("failed to get math grade")
		return nil, fmt.Errorf("failed to get math grade")
	}
	scienceGrade, err := strconv.Atoi(items[2])
	if err != nil {
		fmt.Println("failed to get science grade")
		return nil, fmt.Errorf("failed to get science grade")
	}

	std := &Student{
		name:         items[0],
		mathGrade:    mathGrade,
		scienceGrade: scienceGrade,
	}

	return std, nil
}

func main() {
	fmt.Println("Implementing custom sorting for student struct slice")
	fmt.Println("Enter number of students")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numberOfStudent, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("failed to get number of students")
	}

	students := []Student{}
	fmt.Println("Enter space separated student details i.e. name math_grade science_grade for example dinesh 55 22")
	for i := 0; i < numberOfStudent; i++ {
		fmt.Printf("Enter student details for #%d\n", i+1)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		student, err := getStudentFromStr(scanner.Text())
		if err != nil {
			fmt.Println("failed to get student")
			return
		}
		students = append(students, *student)
	}

	// using built-in sorting function in golang
	sort.SliceStable(students, func(i, j int) bool {
		return students[i].getTotalGrade() > students[j].getTotalGrade()
	})

	fmt.Printf("Sorted students by total grade - %v\n", students)
}
