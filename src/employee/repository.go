package employee

import "context"

type Repository interface {
	HandleEmployeeById(c context.Context, id int) (*Employee, error)
	HandleEmployeeByPosition(c context.Context, position Position) ([]Employee, error)
	InsertEmployee(c context.Context, e *Employee) error
}
