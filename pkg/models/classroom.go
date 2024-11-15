package models

import "time"

type Classroom struct {
	ID       uint      `json:"id"`
	Course   uint      `json:"course"`
	Teacher  uint      `json:"teacher"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}