package service

import "github.com/mct-joken/mitsuami/model/domain"

type ItemService struct {
}

// Create 備品を追加
func (s *ItemService) Create(args CreateItemArgs) error {
	return nil
}

// GetItem 備品を取得
func (s *ItemService) GetItem(id string) domain.Item {
	return domain.Item{}
}

// GetItemList 備品リストを取得
func (s *ItemService) GetItemList() []domain.Item {
	return []domain.Item{}
}

// StartUsing 備品の利用を開始
func (s *ItemService) StartUsing(itemID, userID string) error {
	return nil
}

// EndUsing 備品の使用を終了(返却)
func (s *ItemService) EndUsing(itemID, userID string) error {
	return nil
}

type CreateItemArgs struct {
	Type        string
	Code        string
	Name        string
	Description string
	Image       string
}
