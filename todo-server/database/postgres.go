package database

import (
	"fmt"

	"beards.ly/todo/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgresConnectionString connection string for db
type PostgresConnectionString string

// NewDB given connection string
func NewDB(dsn PostgresConnectionString) *gorm.DB {
	db, err := gorm.Open(postgres.Open(string(dsn)), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	return db
}

// ConnectionString constructs a connection string for postgresql
func ConnectionString(host string, port string, db string, user string, password string) PostgresConnectionString {
	return PostgresConnectionString(fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", host, port, user, db, password))
}

// NewPostgresConnectionString gets the connection string
func NewPostgresConnectionString() PostgresConnectionString {
	cfg := config.GetConfig()
	host := cfg.GetString("POSTGRES_HOST")
	port := cfg.GetString("POSTGRES_PORT")
	database := cfg.GetString("POSTGRES_DB")
	user := cfg.GetString("POSTGRES_USER")
	password := cfg.GetString("POSTGRES_PASSWORD")
	return ConnectionString(host, port, database, user, password)
}
