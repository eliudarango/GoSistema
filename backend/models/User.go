package models

import (

)

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Email string `gorm:"unique;not null"`
}
