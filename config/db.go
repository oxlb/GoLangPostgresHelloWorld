//GetPostgresConnectionString for GORM
package config

import (
	"fmt"
)

// DB CONSTANT
const (
	DBUser     = "root"
	DBPassword = "root"
	DBName     = "root"
	DBHost     = "0.0.0.0"
	DBPort     = "5432"
	DBType     = "postgres"
)

//GetDBType
func GetDBType() string {
	return DBType
}

func GetPostgresConnectionString() string {

	dataBase := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		DBHost,
		DBPort,
		DBUser,
		DBName,
		DBPassword)

	return dataBase
}