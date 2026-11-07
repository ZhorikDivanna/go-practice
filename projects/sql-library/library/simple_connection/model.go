package simple_connection

import "time"

type Book struct {
	ID                int
	Name              string
	Author            string
	Review            *string
	YearOfPublication int
	Completed         bool
	OnTheShelfAt      time.Time
	CompletedAt       *time.Time
}
