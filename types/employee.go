package types

import "fmt"

type Employee struct {
	Id         string
	Name       string
	Age        int
	Department string
	Salary     int
}

func (e *Employee) ToString() string {
	return fmt.Sprintf("Employee Name: %s, Department: %s, Salary: %d, Age: %d", e.Name, e.Department, e.Salary, e.Age)
}
