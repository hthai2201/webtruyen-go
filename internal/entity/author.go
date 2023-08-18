package entity

import "github.com/hthai2201/webtruyen-go/pkg/common"

// Translation -.
type Author struct {
	common.Model
	Slug string `json:"slug" gorm:"column:slug;unique_index"`
	Name string `json:"name" gorm:"column:name"`
}

func (Author) GetTableName() string {
	return "authors"
}
