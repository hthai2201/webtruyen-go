// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

import (
	"github.com/hthai2201/webtruyen-go/pkg/common"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Translation -.
type Story struct {
	common.Model
	Slug        string     `json:"slug" gorm:"column:slug;uniqueIndex"`
	Name        string     `json:"name" gorm:"column:name"`
	Thumbnail   string     `json:"thumbnail" gorm:"column:thumbnail"`
	AuthorId    int        `json:"author_id" gorm:"column:author_id"`
	Author      Author     `json:"author" gorm:"foreignkey:author_id"`
	Categories  []Category `json:"categories" gorm:"many2many:stories_categories;"`
	Source      string     `json:"source" gorm:"column:source"`
	Status      string     `json:"status" gorm:"column:status"`
	Rate        StoryRate  `json:"rate" gorm:"foreignkey:story_id"`
	Description string     `json:"description" gorm:"column:description"`
}

func (Story) TableName() string {
	return "stories"
}

func (e *Story) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Statement.AddClause(clause.OnConflict{
		Columns:   []clause.Column{{Name: "slug"}},
		UpdateAll: true,
	})
	return nil
}
