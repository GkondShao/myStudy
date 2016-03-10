package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Task struct {
	Id      int64  `orm:"pk;auto"` // Unique identifier
	Title   string `orm:"null"`    // Description
	Content string `orm:"null"`
	//User    *User  `orm:"rel(fk)"` //foreign key
	Done bool
	// Is this task done?
}

// NewTask creates a new task given a title, that can't be empty.
func NewTask(title, content string) error {
	if title == "" {
		return fmt.Errorf("empty title")
	}
	task := &Task{0, title, content, false}
	o := orm.NewOrm()
	_, err := o.Insert(task)
	return err
}
