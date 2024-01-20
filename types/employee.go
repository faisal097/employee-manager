package types

import "fmt"

// Employee represent employee details in the system
type Employee struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Department string `json:"department"`
	Salary     int    `json:"salary"`
}

// NewEmployee returns the Employee object
func NewEmployee(id string, name string) *Employee {
	return &Employee{
		ID:   id,
		Name: name,
	}
}

// ToString stringify the Employee object
func (e *Employee) ToString() string {
	return fmt.Sprintf("Employee Name: %s, Department: %s, Salary: %d, Age: %d", e.Name, e.Department, e.Salary, e.Age)
}
