package entity

import "github.com/hthai2201/webtruyen-go/pkg/common"

// Translation -.
type Chapter struct {
	common.Model
	Index   int
	Slug    string `json:"slug" gorm:"column:slug;unique_index"`
	Name    string `json:"name" gorm:"column:name"`
	Content string `json:"content" gorm:"column:content"`
}

func (Chapter) GetTableName() string {
	return "categories"
}
