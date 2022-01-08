package repository_test

import (
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/raihanlh/go-kubernetes/src/entity"
	"github.com/raihanlh/go-kubernetes/src/repository"
	"github.com/stretchr/testify/assert"
)

var b = &entity.Bulletin{
	Author:    "Test Author",
	Content:   "This is a content from test author",
	CreatedAt: time.Now(),
}

func TestFindAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error when opening stub database connection: %v", err)
	}

	rows := sqlmock.NewRows([]string{"author", "content", "created_at"}).AddRow(
		b.Author, b.Content, b.CreatedAt,
	)

	query := `SELECT author, content, created_at FROM bulletins ORDER BY created_at DESC LIMIT 100`

	mock.ExpectQuery(query).WillReturnRows(rows)

	repo := repository.NewPostgresqlBulletinRepository(db)

	articles, err := repo.FindAll()

	assert.NoError(t, err)
	assert.NotNil(t, articles)
}

func TestSave(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error when opening stub database connection: %v", err)
	}

	query := `INSERT INTO bulletins (author, content, created_at) VALUES ($1, $2, $3)`

	mock.ExpectExec(query).WithArgs(b.Author, b.Content, b.CreatedAt).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := repository.NewPostgresqlBulletinRepository(db)

	err = repo.Save(*b)
	assert.Error(t, err)
}
