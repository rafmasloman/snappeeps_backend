package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDatabase() (*gorm.DB, error) {
	dbVariables := EnvDatabase()

	dsn := fmt.Sprintf(`%s:%s@tcp(127.0.0.1:%s)/%s`, dbVariables.DBUSER, dbVariables.DBPASSWORD, dbVariables.DBPORT, dbVariables.DBNAME)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf(`failed to connect database = %w`, err)
	}

	return db, nil
}
