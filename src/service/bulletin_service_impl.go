package service

import (
	"github.com/raihanlh/go-kubernetes/src/entity"
	"github.com/raihanlh/go-kubernetes/src/repository"
)

type BulletinServiceImpl struct {
	bulletinRepo repository.BulletinRepository
}

func NewBulletinService(bulletinRepository repository.BulletinRepository) BulletinService {
	return &BulletinServiceImpl{bulletinRepo: bulletinRepository}
}

func (b *BulletinServiceImpl) GetAll() ([]entity.Bulletin, error) {
	response, err := b.bulletinRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (b *BulletinServiceImpl) InsertOne(bulletin entity.Bulletin) error {
	err := b.bulletinRepo.Save(bulletin)
	if err != nil {
		return err
	}
	return nil
}
