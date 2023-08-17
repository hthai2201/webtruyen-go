package entity

import "github.com/hthai2201/webtruyen-go/pkg/common"

// Translation -.
type StoryList struct {
	common.Model
	Slug        string `json:"slug" gorm:"column:slug;unique_index"`
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
}

func (StoryList) GetTableName() string {
	return "story_list"
}
