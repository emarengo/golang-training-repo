package employee

import "time"

type Employee struct {
	ID          int       `json:"id,omitempty"`
	FullName    string    `json:"full_name,omitempty"`
	Position    Position  `json:"position,omitempty"`
	Salary      float64   `json:"salary,omitempty"`
	Joined      time.Time `json:"joined"`
	OnProbation bool      `json:"on_probation,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}
