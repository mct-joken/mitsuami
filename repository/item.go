package repository

import "github.com/mct-joken/mitsuami/model/domain"

type ItemRepository interface {
	CreateItem(item domain.Item) error
	GetItem() []domain.Item
	GetItemByID(id string) *domain.Item
}
