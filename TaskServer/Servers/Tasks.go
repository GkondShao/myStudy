package Servers

import (
	. "TaskServer/models"
	"database/sql"
	"fmt"
	"log"

	. "TaskServer/confutil"

	_ "github.com/Go-SQL-Driver/MySQL"
)

type TaskManager int

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (t *TaskManager) Query(args Args, reply *Tasks) error {
	c := new(Config)
	c.Anly("config\\app.conf")
	dbu := c.Secs["DB"]["user"]
	dbp := c.Secs["DB"]["pwd"]
	dbn := c.Secs["DB"]["name"]

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8", dbu, dbp, dbn))
	checkErr(err)

	defer db.Close()

	err = db.Ping()
	checkErr(err)

	rows, err := db.Query("SELECT * FROM TASk WHERE Done = ?", args.Done)
	checkErr(err)
	defer rows.Close()
	for rows.Next() {
		var task Task
		err = rows.Scan(&task.Id, &task.Title, &task.Content, &task.Done)
		checkErr(err)
		reply.TaskList = append(reply.TaskList, task)
	}
	return nil
}

func (t *TaskManager) Insert(args Task, reply *int) error {
	c := new(Config)
	c.Anly("config\\app.conf")
	dbu := c.Secs["DB"]["user"]
	dbp := c.Secs["DB"]["pwd"]
	dbn := c.Secs["DB"]["name"]

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8", dbu, dbp, dbn))
	if err != nil {
		log.Fatal(err)
		*reply = 500
		return nil
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		*reply = 500
		return nil
	}

	stmt, err := db.Prepare("INSERT INTO TASK(Title,Content,Done)values(?,?,?)")

	if err != nil {
		log.Fatal(err)
		*reply = 500
		return nil
	}
	defer stmt.Close()
	rs, err := stmt.Exec(args.Title, args.Content, false)
	if err != nil {
		log.Fatal(err)
		*reply = 500
		return nil
	}
	aff, err := rs.RowsAffected()
	if err != nil {
		log.Fatal(err)
		*reply = 500
		return nil
	}

	if aff == 1 {
		*reply = 200
	}
	return nil
}
