package main

import "time"

type User struct {
	ID           int        `gorm:"column:id;primary_key"`
	FirstName    string     `gorm:"column:first_name"`
	LastName     string     `gorm:"column:last_name"`
	Emails       []Email    `gorm:"foreignkey:USERID;constraint:onDelete:CASCADE"` // One-to-many relationship
	PasswordHash string     `gorm:"column:password_hash"`
	LastIP       string     `gorm:"column:last_ip"`
	CreatedAt    *time.Time `gorm:"column:created_at"`
	UpdatedAt    *time.Time `gorm:"column:updated_at"`
}

type Email struct {
	ID      int    `gorm:"column:id;primary_key"`
	Address string `gorm:"column:address;size:256;uniqueIndex"`
	Primary bool   `gorm:"column:primary"`
	UserID  int    `gorm:"column:user_id"`
}
