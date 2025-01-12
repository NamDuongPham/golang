package model

import (
	"errors"
	"my-app/common"
)

var (
	ErrTitleIsBlank = errors.New("title can not be blank")
)

type TodoItem struct {
	common.SQLModel
	Title string `json:"title" gorm:"column:title"`
	// Image string `json:"image"`
	Description string      `json:"description" gorm:"column:description"`
	Status      *ItemStatus `json:"status" gorm:"column:status"`
}
type TodoItemCreate struct {
	Id          int         `json:"-" gorm:"column:id"`
	Title       string      `json:"title" gorm:"column:title"`
	Description string      `json:"description" gorm:"column:description"`
	Status      *ItemStatus `json:"status" gorm:"column:status"`
}
type TodoItemUpdate struct {
	Title       *string `json:"title" gorm:"column:title"`
	Description *string `json:"description" gorm:"column:description"`
	Status      *string `json:"status" gorm:"column:status"`
}

func (TodoItem) TableName() string {
	return "todo_items"
}
func (TodoItemCreate) TableName() string {
	return TodoItem{}.TableName()
}
func (TodoItemUpdate) TableName() string {
	return TodoItem{}.TableName()
}
