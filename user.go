package avitoTech

type User struct {
	ID      uint64  `json:"id" db:"id" binding:"required"`
	Balance float64 `json:"balance" binding:"min=0"`
}

type Id struct {
	ID uint64 `json:"id" db:"id" binding:"required"`
}

type UserAndCount struct {
	ID    uint64  `json:"id" db:"id" binding:"required"`
	Count float64 `json:"count" binding:"min=0"`
}

type SenderRecipientAndCount struct {
	Sender    uint64  `json:"sender" db:"id" binding:"required"`
	Recipient uint64  `json:"recipient" db:"id" binding:"required"`
	Count     float64 `json:"count" binding:"min=0"`
}

type UserServiceAndCount struct {
	UserId      uint64  `json:"userId" binding:"required"`
	ServiceId   uint64  `json:"serviceId" binding:"required"`
	Count       float64 `json:"count" binding:"min=0"`
	Description string  `json:"description" binding:"required"`
}
