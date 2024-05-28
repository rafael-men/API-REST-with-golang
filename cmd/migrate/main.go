package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/rafael-men/rest-api-with-golang/cmd/api"
	"github.com/rafael-men/rest-api-with-golang/config"
	"github.com/rafael-men/rest-api-with-golang/db"
)

func main() {

	cfg := config.EnvConfig

	mysqlCfg := mysql.Config{
		User:                 cfg.DBUser,
		Passwd:               cfg.DBPassword,
		Addr:                 cfg.DBAddress,
		DBName:               cfg.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	dsn := mysqlCfg.FormatDSN()

	storage, err := db.NewMySQLStorage(dsn)
	if err != nil {
		log.Fatalf("failed to connect to %v", err)
	}

	server := api.NewAPIServer(":8080", storage)
	if err := server.Run(); err != nil {
		log.Fatalf("Unexpected Error: %v", err)
	} else {
		log.Println("Back-End running on port 8080")
	}
}
