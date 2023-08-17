package entity

import (
	"github.com/hthai2201/webtruyen-go/pkg/common"
)

type StoryRate struct {
	common.Model
	Value     float32 `json:"value" gorm:"column:value"`
	BestValue float32 `json:"best_value" gorm:"column:best_value"`
	Count     int     `json:"count" gorm:"column:count"`
}

func (StoryRate) GetTableName() string {
	return "story_rates"
}
