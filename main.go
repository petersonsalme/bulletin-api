package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	DbHost     = "db"
	DbUser     = "postgres-dev"
	DbPassword = "mysecretpassword"
	DbName     = "dev"
)

type Bulletin struct {
	Author    string    `json:"author" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

var db *sql.DB

func GetBulletins() ([]Bulletin, error) {
	const q = `SELECT author, content, created_at FROM bulletins ORDER BY created_at DESC LIMIT 100`

	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}

	results := make([]Bulletin, 0)

	for rows.Next() {
		var author, content string
		var createdAt time.Time

		err = rows.Scan(&author, &content, &createdAt)
		if err != nil {
			return nil, err
		}

		results = append(results, Bulletin{author, content, createdAt})
	}

	return results, nil
}

func AddBulletin(b Bulletin) error {
	const q = `INSERT INTO bulletins (author, content, created_at) VALUES ($1, $2, $3)`

	_, err := db.Exec(q, b.Author, b.Content, b.CreatedAt)

	return err
}

func main() {
	var err error

	r := gin.Default()
	r.GET("/board", func(context *gin.Context) {
		results, err := GetBulletins()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"status": "Internal Server Error: " + err.Error()})
			return
		}
		context.JSON(http.StatusOK, results)
	})
	r.POST("/board", func(context *gin.Context) {
		var b Bulletin

		if context.Bind(&b) == nil {
			b.CreatedAt = time.Now()
			if err := AddBulletin(b); err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"status": "Internal Server Error: " + err.Error()})
				return
			}
			context.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", DbHost, DbUser, DbPassword, DbName)
	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var migration *string
	migration, err = Migration()
	if err != nil {
		panic(err)
	}

	_, err = db.Query(*migration)
	if err != nil {
		log.Println("Failed to run migrations:")
		log.Println(err)
		return
	}

	log.Println("running...")
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

func Migration() (*string, error) {
	file, err := ioutil.ReadFile("./migration.sql")
	if err != nil {
		return nil, fmt.Errorf("Failed to read migration.sql\nError: %s", err.Error())
	}

	s := string(file)
	return &s, nil
}
