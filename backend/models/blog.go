package models

import ()


type Blog struct{

	ID    uint   `gorm:"primaryKey"`
	Title  string `gorm:"not null"`
	Content string `gorm:"unique;not null"`

}