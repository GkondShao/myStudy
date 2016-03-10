package main

import (
	"TaskClient/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var outTE, inContent *walk.TextEdit
var inTitle *walk.LineEdit
var task models.Task
var view models.TaskView

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}
func loadData() {
	url, _ := url.Parse("http://127.0.0.1:8080/task")
	resp, err := http.Get(url.String())
	checkErr(err)

	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	defer resp.Body.Close()

	var taskList models.TaskList
	json.Unmarshal(body, &taskList)
	var str string
	strs := "this is your task todo\r\n"

	for _, v := range taskList.Tasks {
		str = v.Title + ":" + v.Content + "\r\n"
		strs += str
	}
	fmt.Println(strs)
	view.Content = strs

}

func sendData(title, content string) {
	task.Title = title
	task.Content = content

	b, err := json.Marshal(task)
	if err != nil {
		log.Fatal("json err : ", err)
	}

	body := bytes.NewBuffer([]byte(b))
	_, err = http.Post("http://127.0.0.1:8080/task", "application/json;charset=utf-8", body)

	checkErr(err)
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
