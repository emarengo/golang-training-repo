package methods

import (
	"context"
	"github.com/stretchr/testify/assert"
	"golang-training-repo/employee"
	"testing"
	"time"
)

func TestCreateEmployee(t *testing.T) {
	employee1 := &employee.Employee{
		FullName:    "Ernesto",
		Position:    employee.CxAgent,
		Salary:      1000,
		Joined:      time.Now(),
		OnProbation: true,
	}
	employee2 := &employee.Employee{
		FullName:    "",
		Position:    employee.CxAgent,
		Salary:      1000,
		Joined:      time.Now(),
		OnProbation: true,
	}
	employee3 := &employee.Employee{
		FullName:    "Ernesto",
		Position:    employee.CxAgent,
		Salary:      0,
		Joined:      time.Now(),
		OnProbation: true,
	}
	type args struct {
		c context.Context
		e *employee.Employee
	}
	testCases := []struct {
		name        string
		args        args
		want        *employee.Employee
		expectedErr string
	}{
		{name: "passed", args: args{c: context.Background(), e: employee1}, want: employee1},
		{name: "full name is empty", args: args{c: context.Background(), e: employee2}, expectedErr: "full name cannot be empty"},
		{name: "salary cannot be 0", args: args{c: context.Background(), e: employee3}, expectedErr: "salary cannot be 0"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := CreateEmployee(tc.args.c, tc.args.e)
			if tc.expectedErr != "" {
				assert.EqualError(t, err, tc.expectedErr)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}

func TestUpdateEmployee(t *testing.T) {
	employee1 := &employee.Employee{
		FullName:    "Ernesto",
		Position:    employee.CxAgent,
		Salary:      1000,
		Joined:      time.Now(),
		OnProbation: true,
	}
	employee2 := &employee.Employee{
		FullName:    "Ernesto",
		Position:    employee.Undefined,
		Salary:      0,
		Joined:      time.Now(),
		OnProbation: true,
	}
	type args struct {
		c context.Context
		e *employee.Employee
	}
	testCases := []struct {
		name        string
		args        args
		want        *employee.Employee
		expectedErr string
	}{
		{name: "passed", args: args{c: context.Background(), e: employee1}, want: employee1},
		{name: "position cannot be undefined", args: args{c: context.Background(), e: employee2}, expectedErr: "position cannot be undefined"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := UpdateEmployee(tc.args.c, tc.args.e)
			if tc.expectedErr != "" {
				assert.EqualError(t, err, tc.expectedErr)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}

func TestGetEmployeeById(t *testing.T) {
	type args struct {
		c  context.Context
		id int
	}
	testCases := []struct {
		name        string
		args        args
		want        *employee.Employee
		expectedErr string
	}{
		{name: "employee not found", args: args{c: context.Background(), id: 2}, want: nil, expectedErr: "employee not found"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetEmployeeById(tc.args.c, tc.args.id)

			if tc.expectedErr != "" {
				assert.EqualError(t, err, tc.expectedErr)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}

func TestGetEmployeesByPosition(t *testing.T) {
	type args struct {
		c        context.Context
		position int
	}
	testCases := []struct {
		name        string
		args        args
		want        []employee.Employee
		expectedErr string
	}{
		{name: "employee not found", args: args{c: context.Background(), position: 2}, want: nil, expectedErr: "employee not found"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetEmployeesByPosition(tc.args.c, tc.args.position)

			if tc.expectedErr != "" {
				assert.EqualError(t, err, tc.expectedErr)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}
