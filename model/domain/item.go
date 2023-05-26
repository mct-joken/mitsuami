package domain

import "time"

type Item struct {
	id          string
	itemType    int
	name        string
	code        string
	description string
	image       string
	createdAt   time.Time
	updatedAt   time.Time
	deletedAt   time.Time
}

func (i *Item) GetID() string {
	return i.id
}

func (i *Item) GetItemType() int {
	return i.itemType
}

func (i *Item) GetName() string {
	return i.name
}

func (i *Item) GetCreatedAt() time.Time {
	return i.createdAt
}

func (i *Item) GetCode() string {
	return i.code
}

func (i *Item) SetCode(code string) {
	i.code = code
}

func (i *Item) GetDescription() string {
	return i.description
}

func (i *Item) SetDescription(description string) {
	i.description = description
}

func (i *Item) GetImage() string {
	return i.image
}

func (i *Item) SetImage(image string) {
	i.image = image
}

func (i *Item) GetUpdatedAt() time.Time {
	return i.updatedAt
}

func (i *Item) SetUpdatedAt(updatedAt time.Time) {
	i.updatedAt = updatedAt
}

func (i *Item) GetDeletedAt() time.Time {
	return i.deletedAt
}

func (i *Item) SetDeletedAt(deletedAt time.Time) {
	i.deletedAt = deletedAt
}

func NewItem(id, name string, itemType int, createdAt time.Time) *Item {
	// id, name, itemType, createdAt は変更不可
	return &Item{
		id:        id,
		name:      name,
		itemType:  itemType,
		createdAt: createdAt,
	}
}
