package domain

import "time"

// Log 貸出記録
type Log struct {
	id        string
	userID    string
	itemID    string
	due       time.Time
	status    int
	createdAt time.Time
	updatedAt time.Time
}

func (l *Log) GetID() string {
	return l.id
}

func (l *Log) GetUserID() string {
	return l.userID
}

func (l *Log) GetItemID() string {
	return l.itemID
}

func (l *Log) CreatedAt() time.Time {
	return l.createdAt
}

func (l *Log) GetDue() time.Time {
	return l.due
}

func (l *Log) SetDue(due time.Time) {
	l.due = due
}

func (l *Log) GetStatus() int {
	return l.status
}

func (l *Log) SetStatus(status int) {
	l.status = status
}

func (l *Log) GetUpdatedAt() time.Time {
	return l.updatedAt
}

func (l *Log) SetUpdatedAt(updatedAt time.Time) {
	l.updatedAt = updatedAt
}

// NewLog 貸出記録オブジェクトを生成
func NewLog(id, userID, itemID string, createdAt time.Time) *Log {
	// id, userID, itemID, createdAtは変更不可
	return &Log{
		id:        id,
		userID:    userID,
		itemID:    itemID,
		createdAt: createdAt,
	}
}
