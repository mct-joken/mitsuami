package service

import (
	"errors"
	"github.com/mct-joken/mitsuami/model/domain"
	"github.com/mct-joken/mitsuami/repository"
	"time"
)

type ItemService struct {
	repository    repository.ItemRepository
	logRepository repository.LogRepository
}

func NewItemService(r repository.ItemRepository) *ItemService {
	return &ItemService{
		repository: r,
	}
}

// Create 備品を追加
func (s *ItemService) Create(args CreateItemArgs) error {
	d := domain.NewItem(args.ID, args.Name, args.ItemType, time.Now())
	d.SetImage(args.Image)
	d.SetCode(args.Code)
	d.SetDescription(args.Description)
	if err := s.repository.CreateItem(*d); err != nil {
		return err
	}
	return nil
}

// GetItem 備品を取得
func (s *ItemService) GetItem(id string) (*domain.Item, error) {
	res := s.repository.GetItemByID(id)
	if res == nil {
		return nil, errors.New("not found")
	}
	return res, nil
}

// GetItemList 備品リストを取得
func (s *ItemService) GetItemList() []domain.Item {
	return s.repository.GetItem()
}

// StartUsing 備品の利用を開始
func (s *ItemService) StartUsing(itemID, userID string) error {
	// ToDo: IDの生成方法を決める
	// そもそもいらないかも
	d := domain.NewLog("", userID, itemID, time.Now())
	if err := s.logRepository.CreateLog(*d); err != nil {
		return err
	}
	return nil
}

// EndUsing 備品の使用を終了(返却)
func (s *ItemService) EndUsing(itemID, userID string) error {
	logs := s.logRepository.GetLogByIDs(itemID, userID)

	log := (*domain.Log)(nil)
	for _, v := range logs {
		// 貸出中: 0, 返却済み: 1
		if v.GetStatus() == 0 {
			log = &v
		}
	}
	if log == nil {
		return errors.New("not found")
	}

	log.SetStatus(1)
	log.SetUpdatedAt(time.Now())

	if err := s.logRepository.UpdateLog(*log); err != nil {
		return err
	}

	return nil
}

type CreateItemArgs struct {
	ID          string
	ItemType    int
	Code        string
	Name        string
	Description string
	Image       string
}
