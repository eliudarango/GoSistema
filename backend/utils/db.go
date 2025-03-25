package utils

import (
	"gosistema/backend/database"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func IniciarConexion() *gorm.DB {
	dsn := database.GetDBConexion()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Error en conexion",err)
	}
	return db
}
