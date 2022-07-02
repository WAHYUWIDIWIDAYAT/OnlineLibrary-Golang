package models

// add cart with session and inner join 

type Cart struct {
	ID uint `gorm:"primary_key" json:"id"`
	UserID uint `json:"user_id"`
	BookID uint `json:"book_id"`
	Quantity uint `json:"quantity"`
}
