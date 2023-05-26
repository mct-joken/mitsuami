package domain

import "time"

type User struct {
	id          string
	name        string
	displayName string
	createdAt   time.Time
	updatedAt   time.Time
	deletedAt   time.Time
}

func (u *User) GetID() string {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetDisplayName() string {
	return u.displayName
}

func (u *User) SetDisplayName(displayName string) {
	u.displayName = displayName
}

func (u *User) GetCreatedAt() time.Time {
	return u.createdAt
}

func (u *User) GetUpdatedAt() time.Time {
	return u.updatedAt
}

func (u *User) SetUpdatedAt(updatedAt time.Time) {
	u.updatedAt = updatedAt
}

func (u *User) GetDeletedAt() time.Time {
	return u.deletedAt
}

func (u *User) SetDeletedAt(deletedAt time.Time) {
	u.deletedAt = deletedAt
}

// NewUser ユーザーオブジェクトを生成
func NewUser(id, name string, createdAt time.Time) *User {
	// id, name, createdAt は変更不可
	return &User{
		id:        id,
		name:      name,
		createdAt: createdAt,
	}
}
