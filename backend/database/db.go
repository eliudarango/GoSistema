
package database

import "fmt"

const (
	DBHost     = "serverdb"
	DBPost     = "3306"
	DBUser     = "user"
	DBPassword = "password"
	DBName     = "sistema"
)

func GetDBConexion() string {

	conexion := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",

		DBUser,
		DBPassword,
		DBHost,
		DBPost,
		DBName)
	return conexion
}
