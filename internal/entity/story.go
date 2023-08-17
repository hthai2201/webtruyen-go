// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

import "github.com/hthai2201/webtruyen-go/pkg/common"

// Translation -.
type Story struct {
	common.Model
	Slug        string `json:"slug" gorm:"column:slug;unique_index"`
	Title       string
	Thumbnail   string
	Author      string
	Categories  []Category `gorm:"many2many:story_categories;"`
	Source      string
	Status      string
	Rate        StoryRate
	Description string
}

func (Story) GetTableName() string {
	return "history"
}
