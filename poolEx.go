package main

import (
	"container/list"
	"log"
	"reflect"
	"sync"
	"time"
)

type Pool struct {
	Num     int
	Pool    chan int
	Timeout chan bool
}

type Tasks struct {
	List   *list.List
	result chan interface{}
	Mutex  sync.Mutex
	Status bool
}

func (task *Tasks) Init(resNum int) {
	task.List = list.New()
	task.result = make(chan interface{}, resNum)
	task.Status = false
}

func (pool *Pool) Init(num int) {
	pool.Num = num
	pool.Pool = make(chan int, num)
	for i := 0; i < num; i++ {
		pool.Pool <- i
	}
	pool.Timeout = make(chan bool)
}

func (tasks *Tasks) Add(arg interface{}) {
	tasks.List.PushBack(arg)
	tasks.Status = true
}

func (work Work) Do(arg interface{}, tasks *Tasks, pool *Pool) {
	value := reflect.ValueOf(arg).Int()
	//log.Println(value)
	time.Sleep(time.Second * 2)
	tasks.result <- value
	pool.Pool <- 1
}

type Process struct{}
type Work struct{}

func (process Process) Process(tasks *Tasks, pool *Pool) {
	for {
		length := tasks.List.Len()
		if length == 0 && len(pool.Pool) == pool.Num {
			pool.Timeout <- true
			tasks.Status = false
		} else if length > 0 {
			_, ok := <-pool.Pool
			if ok {
				work := Work{}
				tasks.Mutex.Lock()
				value := tasks.List.Front()
				tasks.List.Remove(value)
				go work.Do(value.Value, tasks, pool)
				tasks.Mutex.Unlock()
			}
		} else if length == 0 && len(pool.Pool) < pool.Num {
			log.Println("任务分配结束，等待线程作业完成，回收")
		}
		time.Sleep(time.Second)
	}
}

func (tasks *Tasks) Listen(process Process, pool *Pool) {
	log.Println("Listen start...")
	for {
		if tasks.Status {
			process.Process(tasks, pool)
		} else {
			log.Println("wait...")
		}
	}
}

func main() {
	tasks := &Tasks{}
	pool := &Pool{}
	tasks.Init(5)
	pool.Init(5)
	process := Process{}
	go tasks.Listen(process, pool)
	for i := 0; i < 10; i++ {
		tasks.Add(i)
	}
	for {
		select {
		case v := <-tasks.result:
			log.Println(v)
		case <-pool.Timeout:
			log.Println("timeout")
		}
	}
}
