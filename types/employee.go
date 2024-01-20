package types

import "fmt"

type Employee struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Department string `json:"department"`
	Salary     int    `json:"salary"`
}

func NewEmployee(id string, name string) *Employee {
	return &Employee{
		Id:   id,
		Name: name,
	}
}

func (e *Employee) ToString() string {
	return fmt.Sprintf("Employee Name: %s, Department: %s, Salary: %d, Age: %d", e.Name, e.Department, e.Salary, e.Age)
}
