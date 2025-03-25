package model

import (

)

type Usuario struct {
	ID    uint   `gorm:"primaryKey"`
	Nombre  string `gorm:"not null"`
	Correo string `gorm:"unique;not null"`
}
