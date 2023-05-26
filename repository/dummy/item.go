package dummy

import (
	"github.com/mct-joken/mitsuami/model/domain"
)

type ItemRepository struct {
	data []domain.Item
}

func (i *ItemRepository) CreateItem(item domain.Item) error {
	i.data = append(i.data, item)
	return nil
}

func (i *ItemRepository) GetItem() []domain.Item {
	return i.data
}

func (i *ItemRepository) GetItemByID(id string) *domain.Item {
	for _, v := range i.data {
		if v.GetID() == id {
			return &v
		}
	}

	return nil
}

func NewItemRepository(data []domain.Item) *ItemRepository {
	return &ItemRepository{data: data}
}
