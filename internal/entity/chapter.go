package entity

import (
	"github.com/hthai2201/webtruyen-go/pkg/common"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Translation -.
type Chapter struct {
	common.Model
	CID     string `json:"cid" gorm:"column:cid;uniqueIndex"`
	StoryID int    `json:"story_id" gorm:"column:story_id"`
	Slug    string `json:"slug" gorm:"column:slug"`
	Name    string `json:"name" gorm:"column:name"`
	Content string `json:"content" gorm:"column:content"`
}

func (Chapter) TableName() string {
	return "chapters"
}

func (e *Chapter) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Statement.AddClause(clause.OnConflict{
		Columns:   []clause.Column{{Name: "cid"}},
		UpdateAll: true,
	})
	return nil
}
