//GetPostgresConnectionString for GORM
package config

import (
	"fmt"
)

// DB CONSTANT
// For Config File Best Practise We have written article please check the following link
// https://medium.com/@onexlab.io/golang-config-file-best-practise-d27d6a97a65a
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
