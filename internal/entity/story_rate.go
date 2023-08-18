package entity

import (
	"github.com/hthai2201/webtruyen-go/pkg/common"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type StoryRate struct {
	common.Model
	StoryID   int     `json:"story_id" gorm:"column:story_id;uniqueIndex"`
	Value     float32 `json:"value" gorm:"column:value"`
	BestValue float32 `json:"best_value" gorm:"column:best_value"`
	Count     int     `json:"count" gorm:"column:count"`
}

func (StoryRate) TableName() string {
	return "story_rates"
}
func (e *StoryRate) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Statement.AddClause(clause.OnConflict{
		Columns:   []clause.Column{{Name: "story_id"}},
		UpdateAll: true,
	})
	return nil
}
