package model

import "time"

// User user
type User struct {
	ID        int
	Name      string
	CreatedAt time.Time
}
