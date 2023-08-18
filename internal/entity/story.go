// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

import "github.com/hthai2201/webtruyen-go/pkg/common"

// Translation -.
type Story struct {
	common.Model
	Slug        string     `json:"slug" gorm:"column:slug;unique_index"`
	Name        string     `json:"name" gorm:"column:name"`
	Thumbnail   string     `json:"thumbnail" gorm:"column:thumbnail"`
	Author      Author     `json:"author"`
	Categories  []Category `gorm:"many2many:stories_categories;"`
	Source      string     `json:"source" gorm:"column:source"`
	Status      string     `json:"status" gorm:"column:status"`
	Rate        StoryRate  `json:"rate"`
	Description string     `json:"description" gorm:"column:description"`
}

func (Story) GetTableName() string {
	return "stories"
}
