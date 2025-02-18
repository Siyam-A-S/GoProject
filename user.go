package main

type User struct {
	id        int
	firstName string
	lasName   string
	emails    []Email
}

type Email struct {
	ID      int    `gorm:"column:id;primary_key"`
	Address string `gorm:"column:address;size:256;uniqueIndex"`
	Primary bool   `gorm:"column:primary"`
	UserID  int    `gorm:"column:user_id"`
}
