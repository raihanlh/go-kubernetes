package service

import "github.com/raihanlh/go-kubernetes/src/entity"

type BulletinService interface {
	GetAll() ([]entity.Bulletin, error)
	InsertOne(bulletin entity.Bulletin) error
}
