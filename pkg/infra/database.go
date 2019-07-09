package infra

import (
	"database/sql"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"os"
)

func NewDatabaseConnection() *sql.DB {
	var name = os.Getenv("DB_NAME")
	var user = os.Getenv("DB_USER")
	var pass = os.Getenv("DB_PASS")
	var host = os.Getenv("DB_HOST")
	dbConnectionString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",user,pass,host,name)
	logger.Info("Connection data %s",dbConnectionString)
	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		logger.Error(err)
	}
	return db
}
