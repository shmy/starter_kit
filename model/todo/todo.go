package todo

import (
	"starter_kit/model"
)

type Todo struct {
	*model.Model
	Title   string `json:"title,omitempty" gorm:"column:title"`
	Content string `json:"content,omitempty" gorm:"column:content"`
}

func (Todo) TableName() string {
	return "t_todos"
}
