package students

import (
	"errors"
	"fmt"
	"io"
	"sort"
)

type Student struct {
	Name  string
	Age   int
	Grade float64
}

type StudentList struct {
	students []Student
}

func NewStudent(name string, age int, grade float64) (*Student, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	if age < 1 || age > 99 {
		return nil, errors.New("age must be between 1 and 99")
	}
	if grade < 0 || grade > 20 {
		return nil, errors.New("grade must be between 0 and 20")
	}
	return &Student{Name: name, Age: age, Grade: grade}, nil
}

func (sl *StudentList) AddStudents(students ...Student) {
	sl.students = append(sl.students, students...)
}

func (sl *StudentList) RemoveStudent(name string) {
	newStudents := make([]Student, 0, len(sl.students))
	for _, s := range sl.students {
		if s.Name != name {
			newStudents = append(newStudents, s)
		}
	}
	sl.students = newStudents
}

func (sl *StudentList) Sort() StudentList {
	newList := make([]Student, len(sl.students))
	copy(newList, sl.students)
	sort.SliceStable(newList, func(i, j int) bool {
		return newList[i].Grade > newList[j].Grade
	})
	return StudentList{students: newList}
}

func (sl *StudentList) Print(out io.Writer) {
	for _, s := range sl.students {
		fmt.Fprintf(out, "%s (%d): %.1f\n", s.Name, s.Age, s.Grade)
	}
}
