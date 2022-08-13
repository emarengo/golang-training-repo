package employee

import (
	"time"
)

type Employee struct {
	ID          int       `json:"id"`
	FullName    string    `json:"fullName"`
	Position    Position  `json:"position"`
	Salary      float64   `json:"salary"`
	Joined      time.Time `json:"joined"`
	OnProbation bool      `json:"onProbation"`
	CreatedAt   time.Time `json:"createdAt"`
}
