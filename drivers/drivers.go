package drivers

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-rest-skeleton/utils"
	"log"
)

func ConnectDB() *sql.DB {

	driver := utils.GetEnv("dbDriver", "mysql")
	username := utils.GetEnv("dbUser", "")
	password := utils.GetEnv("dbPwd", "")
	host := utils.GetEnv("dbHost", "127.0.0.1")
	port := utils.GetEnv("dbPort", "3306")
	database := utils.GetEnv("dbName", "")

	if username == "" || password == "" || database == "" {
		log.Fatal("Invalid database configuration. Check .env file.")
	}

	dbUrl := fmt.Sprintf(username + ":" + password + "@(" + host + ":" + port + ")/" + database)

	db, err := sql.Open(driver, dbUrl)
	utils.LogFatal(err)

	return db
}
