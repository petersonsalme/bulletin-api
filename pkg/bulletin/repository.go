package bulletin

import (
	"database/sql"
	"time"

	"github.com/petersonsalme/bulletin-api/pkg/entity"
)

//PostgresRepository postgres repo
type PostgresRepository struct {
	db *sql.DB
}

//NewPostgresRepository create new repository
func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) GetAll() ([]entity.Bulletin, error) {
	const q = `
		SELECT author, content, created_at 
		FROM bulletins 
		ORDER BY created_at DESC 
		LIMIT 100
	`
	rows, err := r.db.Query(q)
	if err != nil {
		return nil, err
	}

	results := make([]entity.Bulletin, 0)

	for rows.Next() {
		var author, content string
		var createdAt time.Time

		err = rows.Scan(&author, &content, &createdAt)
		if err != nil {
			return nil, err
		}

		results = append(results, entity.Bulletin{
			Author:    author,
			Content:   content,
			CreatedAt: createdAt,
		})
	}

	return results, nil
}

func (r *PostgresRepository) Add(b entity.Bulletin) error {
	const q = `
		INSERT INTO bulletins (author, content, created_at) 
		VALUES ($1, $2, $3)
	`

	_, err := r.db.Exec(q, b.Author, b.Content, b.CreatedAt)

	return err
}
