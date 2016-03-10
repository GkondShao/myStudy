package models

type TaskList struct {
	Tasks []Task
}

type Task struct {
	Id      int64
	Title   string
	Content string
	//User    *User  `orm:"rel(fk)"` //foreign key
	Done bool
	// Is this task done?
}

type TaskView struct {
	Content string
}
