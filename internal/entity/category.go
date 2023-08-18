package entity

import (
	"github.com/hthai2201/webtruyen-go/pkg/common"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Translation -.
type Category struct {
	common.Model
	Slug        string `json:"slug" gorm:"column:slug;uniqueIndex"`
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
}

func (Category) TableName() string {
	return "categories"
}
func (e *Category) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Statement.AddClause(clause.OnConflict{
		Columns:   []clause.Column{{Name: "slug"}},
		UpdateAll: true,
	})
	return nil
}
