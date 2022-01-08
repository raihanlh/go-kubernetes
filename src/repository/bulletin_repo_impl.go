package repository

import (
	"database/sql"
	"time"

	"github.com/raihanlh/go-kubernetes/src/entity"
)

type postgresqlBulletinRepo struct {
	DB *sql.DB
}

func NewPostgresqlBulletinRepository(db *sql.DB) BulletinRepository {
	return &postgresqlBulletinRepo{
		DB: db,
	}
}

func (p *postgresqlBulletinRepo) FindAll() ([]entity.Bulletin, error) {
	const query = `SELECT author, content, created_at FROM bulletins ORDER BY created_at DESC LIMIT 100`
	rows, err := p.DB.Query(query)
	if err != nil {
		return nil, err
	}

	results := make([]entity.Bulletin, 0)

	for rows.Next() {
		var author string
		var content string
		var created_at time.Time
		err = rows.Scan(&author, &content, &created_at)
		if err != nil {
			return nil, err
		}
		results = append(results, entity.Bulletin{Author: author, Content: content, CreatedAt: created_at})
	}

	return results, nil
}

func (p *postgresqlBulletinRepo) Save(bulletin entity.Bulletin) error {
	const query = `INSERT INTO bulletins (author, content, created_at) VALUES ($1, $2, $3)`

	_, err := p.DB.Exec(query, bulletin.Author, bulletin.Content, bulletin.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
