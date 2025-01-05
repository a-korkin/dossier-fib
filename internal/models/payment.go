package models

import "time"

type Payment struct {
	ID       uint32    `json:"id"`
	PersonID uint32    `json:"person_id"`
	Sum      float32   `json:"sum"`
	Created  time.Time `json:"created"`
}

type PaymentDTO struct {
	Sum float32 `json:"sum"`
}
