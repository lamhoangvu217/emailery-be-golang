package models

import "time"

type Email struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Address   string    `gorm:"type:varchar(255);uniqueIndex" json:"address"` // Specify type as varchar
	CreatedAt time.Time `gorm:"type:datetime" json:"created_at"`
}

type Message struct {
	ID         string    `gorm:"primaryKey" json:"id"`
	EmailID    string    `gorm:"index;not null" json:"email_id"`
	Sender     string    `json:"sender"`
	Subject    string    `json:"subject"`
	Body       string    `json:"body"`
	ReceivedAt time.Time `json:"received_at"`

	Email Email `gorm:"constraint:OnDelete:CASCADE;"` // Foreign key relation
}
