package bulletin

import "github.com/petersonsalme/bulletin-api/pkg/entity"

type BulletinService struct {
	repo *PostgresRepository
}

func NewBulletinService(r *PostgresRepository) *BulletinService {
	return &BulletinService{
		repo: r,
	}
}

// GetAll queries all bulletins
func (s *BulletinService) GetAll() ([]entity.Bulletin, error) {
	return s.GetAll()
}

// Add adds a new Bulletin to database
func (s *BulletinService) Add(b entity.Bulletin) error {
	return s.Add(b)
}
