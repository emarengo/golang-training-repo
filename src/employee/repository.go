package employee

import "context"

type Repository interface {
	InsertEmployee(ctx context.Context, e *Employee) error
}
