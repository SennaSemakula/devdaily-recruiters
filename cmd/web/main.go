package main

import (
	"database/sql"
	"fmt"
	"github.com/SennaSemakula/recruit-backend/models"
	_ "github.com/lib/pq"
	"log"
)

type DBConfig struct {
	host   string
	port   int
	user   string
	dbName string
}

type App struct {
	db *sql.DB
}

func main() {
	if err := run; err != nil {
		log.Fatal(err)
	}
}

func run() error {
	config := DBConfig{"localhost", 5432, "devdaily", "devdailydb"}

	db, err := startDB(&config)

	if err != nil {
		return fmt.Errorf("Unable to start database: %v", err)
	}
	users, err := db.GetUsers()

	fmt.Println(users)

	if err != nil {
		return fmt.Errorf("unable to get users: %v", err)
	}

	return nil

}

func startDB(c *DBConfig) (*sql.DB, error) {
	psConfig := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		c.host, c.port, c.user, c.dbName)

	db, err := sql.Open("postgres", psConfig)

	if err != nil {
		return nil, fmt.Errorf("database config: %v", err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v", err)
	}

	log.Printf("successfully connected to database %s", c.dbName)
	return db, nil
}
