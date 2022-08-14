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
