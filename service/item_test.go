package service

import (
	"github.com/mct-joken/mitsuami/model/domain"
	"github.com/mct-joken/mitsuami/repository/dummy"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestItemService_EndUsing(t *testing.T) {
	s := NewItemService(
		dummy.NewItemRepository([]domain.Item{
			*domain.NewItem("123", "test", 0, time.Now()),
		}),
		dummy.NewLogRepository([]domain.Log{
			*domain.NewLog("2222", "110", "123", time.Now()),
		}),
	)

	err := s.EndUsing("123", "110")
	assert.Equal(t, nil, err)

	err = s.EndUsing("123", "110")
	assert.NotEqual(t, nil, err)
}
