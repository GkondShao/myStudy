package main

import (
	"taskList/controllers"

	"github.com/astaxie/beego"
)

func main() {

	beego.Router("/", &controllers.MainController{})
	beego.Router("/task/", &controllers.TaskController{}, "get:ListTasks;post:NewTask")
	beego.Run()
}
