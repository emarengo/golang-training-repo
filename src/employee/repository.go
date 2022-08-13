package employee

import "context"

type Repository interface {
	HandleEmployee(c context.Context, id int) (*Employee, error)
	HandleEmployees(c context.Context, pos Position) ([]Employee, error)
	InsertEmployee(c context.Context, e *Employee) error
}
