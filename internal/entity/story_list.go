package entity

import (
	"github.com/hthai2201/webtruyen-go/pkg/common"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Translation -.
type StoryList struct {
	common.Model
	Slug        string `json:"slug" gorm:"column:slug;uniqueIndex"`
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
}

func (StoryList) TableName() string {
	return "story_list"
}

func (e *StoryList) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Statement.AddClause(clause.OnConflict{
		Columns:   []clause.Column{{Name: "slug"}},
		UpdateAll: true,
	})
	return nil
}
