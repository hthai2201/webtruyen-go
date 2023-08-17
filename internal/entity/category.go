package entity

import "github.com/hthai2201/webtruyen-go/pkg/common"

// Translation -.
type Category struct {
	common.Model
	Slug        string `json:"slug" gorm:"column:slug;unique_index"`
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
}

func (Category) GetTableName() string {
	return "categories"
}
