package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
	. "taskRPCClient/models"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var outTE, inContent *walk.TextEdit
var inTitle *walk.LineEdit
var task Task
var view TaskView

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}
func loadData() {
	client, err := jsonrpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	args := Args{false}
	reply := new(Tasks)

	err = client.Call("TaskManager.Query", args, reply)
	if err != nil {
		log.Fatal(err)
	}

	var str string
	strs := "this is your task todo\r\n"

	for _, v := range reply.TaskList {
		str = v.Title + ":" + v.Content + "\r\n"
		strs += str
	}
	view.Content = strs

}

func sendData(title, content string) {
	client, err := jsonrpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	var re int
	task := Task{Title: inTitle.Text(), Content: inContent.Text()}
	err = client.Call("TaskManager.Insert", task, &re)
	fmt.Println(re)
	loadData()

}
func main() {
	loadData()
	MainWindow{
		Title:   "TaskList",
		MinSize: Size{400, 300},
		Layout:  VBox{},
		Children: []Widget{
			TextEdit{AssignTo: &outTE, ReadOnly: true, Text: view.Content, ColumnSpan: 3},
			VSplitter{
				Children: []Widget{
					LineEdit{AssignTo: &inTitle, Text: "Title"},
					TextEdit{AssignTo: &inContent, Text: "Content", ColumnSpan: 2},
				},
			},
			PushButton{
				Text: "AddTask",
				OnClicked: func() {
					sendData(inTitle.Text(), inContent.Text())
					outTE.SetText(view.Content)
				},
			},
		},
	}.Run()

}
