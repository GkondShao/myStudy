package models

type Tasks struct {
	TaskList []Task
}

type Task struct {
	Id      int64
	Title   string
	Content string
	Done    bool
}

type Args struct {
	Done bool
}

type TaskView struct {
	Content string
}
