package config

import (
	"fmt"
	"os"
)

type EnvDB struct {
	DBNAME     string
	DBPORT     string
	DBPASSWORD string
	DBUSER     string
}

func EnvDatabase() *EnvDB {
	dbName := os.Getenv("DBNAME")
	dbPort := os.Getenv("DBPORT")
	dbUser := os.Getenv("DBUSER")
	dbPassword := os.Getenv("DBPASSWORD")

	dbEnv := EnvDB{
		DBNAME:     dbName,
		DBPORT:     dbPort,
		DBUSER:     dbUser,
		DBPASSWORD: dbPassword,
	}

	fmt.Println("db name = ", dbName)

	return &dbEnv
}
