package entity

import (
	"github.com/hthai2201/webtruyen-go/pkg/common"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Translation -.
type Chapter struct {
	common.Model
	Index   int    `json:"index" gorm:"column:slug;index"`
	Slug    string `json:"slug" gorm:"column:slug;uniqueIndex"`
	Name    string `json:"name" gorm:"column:name"`
	Content string `json:"content" gorm:"column:content"`
}

func (Chapter) TableName() string {
	return "chapters"
}

func (e *Chapter) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Statement.AddClause(clause.OnConflict{
		Columns:   []clause.Column{{Name: "slug"}},
		UpdateAll: true,
	})
	return nil
}
