// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

import "github.com/hthai2201/webtruyen-go/pkg/common"

// Translation -.
type Translation struct {
	common.Model
	Source      string `json:"source"       example:"auto"`
	Destination string `json:"destination"  example:"en"`
	Original    string `json:"original"     example:"текст для перевода"`
	Translation string `json:"translation"  example:"text for translation"`
}

func (Translation) TableName() string {
	return "history"
}
