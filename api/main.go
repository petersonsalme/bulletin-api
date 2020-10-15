package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/petersonsalme/bulletin-api/pkg/bulletin"

	"github.com/gin-gonic/gin"
	"github.com/petersonsalme/bulletin-api/api/handler"

	_ "github.com/lib/pq"
)

const (
	// DbHost database hostname
	DbHost = "db"
	// DbUser database user
	DbUser = "postgres-dev"
	// DbPassword database password
	DbPassword = "mysecretpassword"
	// DbName database name
	DbName = "dev"
)

func main() {
	db := Connect()

	RunMigrations(db)

	repo := bulletin.NewPostgresRepository(db)
	service := bulletin.NewBulletinService(repo)

	r := gin.Default()

	handler.BulletinRouteCfg(r, service)

	log.Println("running...")
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

// Connect opens a new database connection
func Connect() *sql.DB {
	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", DbHost, DbUser, DbPassword, DbName)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}
	return db
}

// RunMigrations runs all migration files
func RunMigrations(db *sql.DB) error {
	filesInfo, err := ioutil.ReadDir("./../migrations")
	if err != nil {
		return fmt.Errorf("Failed to read migrations dir.\nError: %v", err.Error())
	}

	for _, file := range filesInfo {
		migration, err := ioutil.ReadFile(file.Name())
		if err != nil {
			return fmt.Errorf("Failed to read file \"%s\".\nError: %v", file.Name(), err.Error())
		}

		_, err = db.Query(string(migration))
		if err != nil {
			err = fmt.Errorf("Failed to run migration \"%s\".\nError: %v", file.Name(), err.Error())
			log.Println(err)
			return err
		}

	}

	return nil
}
