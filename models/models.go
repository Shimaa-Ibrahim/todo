package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique; not null;default:null"`
	Password string `gorm:"not null;default:null"`
}

type Task struct {
	gorm.Model
	UserID      uint   `gorm:"TYPE:integer REFERENCES users; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Text        string `gorm:"not null;default:null"`
	IsCompleted bool   `gorm:"default:false"`
}

func (c *User) TableName() string {
	return "todo.users"
}
func (c *Task) TableName() string {
	return "todo.tasks"
}
