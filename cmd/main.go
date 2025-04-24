// This file initializes Logging, Database connection and starts the API server.
package main

import (
	"cema_backend/cmd/app"
	"cema_backend/config"
	"cema_backend/db"
	"cema_backend/logging"
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	logging.Initialize()

	// initializes db using credentials from the config file (.env)
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUSER,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	// if unable to connect the db, it logs the error and terminates
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	// initializes the API server using the db connection and the port from the config file
	server := app.NewAPIServer(":"+config.Envs.Port, db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

// initStorage checks if the database connection is alive and logs the status
func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	log.Println("DB: Online")
}
