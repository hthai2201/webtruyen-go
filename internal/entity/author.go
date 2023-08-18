package entity

import (
	"github.com/hthai2201/webtruyen-go/pkg/common"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Translation -.
type Author struct {
	common.Model
	Slug string `json:"slug" gorm:"column:slug;uniqueIndex"`
	Name string `json:"name" gorm:"column:name"`
}

func (Author) TableName() string {
	return "authors"
}

func (e *Author) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Statement.AddClause(clause.OnConflict{
		Columns:   []clause.Column{{Name: "slug"}},
		UpdateAll: true,
	})
	return nil
}
