package types_test

import (
	"reflect"
	"testing"

	"github.com/faisal097/employee-manager/types"
)

func TestNewEmployee(t *testing.T) {
	tests := []struct {
		testName string
		id       string
		name     string
		output   *types.Employee
	}{
		{
			testName: "Test 1",
			id:       "emp00001",
			name:     "Faisal",
			output: &types.Employee{
				ID:   "emp00001",
				Name: "Faisal",
			},
		},
	}

	for _, test := range tests {
		got := types.NewEmployee(test.id, test.name)
		if !reflect.DeepEqual(test.output, got) {
			t.Errorf("Test: %s failed, Not matching.", test.testName)
		}
	}
}

func TestToString(t *testing.T) {
	tests := []struct {
		testName string
		id       string
		name     string
		output   string
	}{
		{
			testName: "Test 1",
			id:       "emp00001",
			name:     "Faisal",
			output:   "Employee Name: Faisal, Department: , Salary: 0, Age: 0",
		},
		{
			testName: "Test 2",
			id:       "emp00002",
			name:     "Saliq",
			output:   "Employee Name: Saliq, Department: , Salary: 0, Age: 0",
		},
	}

	for _, test := range tests {
		emp := types.NewEmployee(test.id, test.name)
		got := emp.ToString()
		if got != test.output {
			t.Errorf("Test: %s failed, Expected: %s, Got: %s", test.testName, test.output, got)
		}
	}
}
