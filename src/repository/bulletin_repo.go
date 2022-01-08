package repository

import (
	"github.com/raihanlh/go-kubernetes/src/entity"
)

type BulletinRepository interface {
	FindAll() ([]entity.Bulletin, error)
	Save(bulletin entity.Bulletin) error
}
