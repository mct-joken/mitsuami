package dummy

import (
	"errors"
	"github.com/mct-joken/mitsuami/model/domain"
)

type LogRepository struct {
	data []domain.Log
}

func (l LogRepository) CreateLog(log domain.Log) error {
	l.data = append(l.data, log)
	return nil
}

func (l LogRepository) GetLogByID(id string) *domain.Log {
	for _, v := range l.data {
		if v.GetID() == id {
			return &v
		}
	}
	return nil
}

func (l LogRepository) GetLogByIDs(itemID, userID string) []domain.Log {
	res := make([]domain.Log, 0)
	for _, v := range l.data {
		if v.GetItemID() == itemID && v.GetUserID() == userID {
			res = append(res, v)
		}
	}

	return res
}

func (l LogRepository) GetLogByItemID(id string) []domain.Log {
	res := make([]domain.Log, 0)
	for _, v := range l.data {
		if v.GetItemID() == id {
			res = append(res, v)
		}
	}
	return res
}

func (l LogRepository) UpdateLog(log domain.Log) error {
	for _, v := range l.data {
		if v == log {
			v = log
			return nil
		}
	}

	return errors.New("not found")
}

func NewLogRepository(data []domain.Log) *LogRepository {
	return &LogRepository{data: data}
}
