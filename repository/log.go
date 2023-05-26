package repository

import "github.com/mct-joken/mitsuami/model/domain"

type LogRepository interface {
	CreateLog(log domain.Log) error
	GetLogByID(id string) *domain.Log
	GetLogByIDs(itemID, userID string) []domain.Log
	GetLogByItemID(id string) []domain.Log
	UpdateLog(log domain.Log) error
}
