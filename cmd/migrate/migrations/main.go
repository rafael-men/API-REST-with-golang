package migrations

import (
	"database/sql"
	"log"

	mysqlCfg "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rafael-men/rest-api-with-golang/config"
)

func main() {
	// Carregar as configurações do ambiente
	cfg := config.EnvConfig

	// Configurar o driver MySQL
	mysqlCfg := mysqlCfg.Config{
		User:                 cfg.DBUser,
		Passwd:               cfg.DBPassword,
		Addr:                 cfg.DBAddress,
		DBName:               cfg.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	log.Printf("Configurações do MySQL: %+v", mysqlCfg)

	db, err := sql.Open("mysql", mysqlCfg.FormatDSN())
	if err != nil {
		log.Fatalf("Erro ao abrir conexão com o banco de dados: %v", err)
	}
	defer db.Close()

	// Obter o driver para migração
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("Erro ao obter driver MySQL para migração: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://cmd/migrate/migrations", "mysql", driver)
	if err != nil {
		log.Fatal(err)
	}
}
