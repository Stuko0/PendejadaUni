package models

import "time"

type Student struct {
	ID          uint      `json:"id"`
	FirstName   string    `json:"firstName"`
	Lastname    string    `json:"lastname"`
	ClassroomID uint      `json:"classroom"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
}
