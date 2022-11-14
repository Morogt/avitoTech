package avitoTech

import (
	"time"
)

type Service struct {
	ID    uint64 `json:"id" db:"id" binding:"required"`
	Title string `json:"title" binding:"required"`
}

type Report struct {
	ID     uint64  `json:"id" db:"id" binding:"required"`
	Amount float64 `json:"title" binding:"required"`
}

type History struct {
	OrderId     uint64
	ServiceId   uint32
	Amount      float64
	Description string
	Refill      bool
	Time        time.Time
}
