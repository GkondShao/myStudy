package taskTimer

//  practice of time.NewTicker
import (
	"errors"
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

const (
	Stop    = "stop"
	Pause   = "pause"
	Running = "Running"
)

var tasks *taskPool

func init() {
	tasks = &taskPool{pool: make(map[string]*Task), isOk: true}
}

type taskPool struct {
	pool map[string]*Task
	m    sync.Mutex
	isOk bool
}

type Task struct {
	Name      string
	Status    string
	Duration  time.Duration
	Content   func() error
	m         sync.Mutex
	CreatTime time.Time
	Chan      chan string
}

//if the name is exist in the pool .it will return a error of "the task is exist"
//if you want to replace the func of the task .please use the function replaceTask(name,func() error)
func AddToPool(name string, t *Task) error {
	tasks.m.Lock()
	defer tasks.m.Unlock()

	if tasks.pool[name] != nil {
		return errors.New(fmt.Sprintf("the task namede %s is exist", name))
	}

	tasks.pool[name] = t

	return nil
}

//todo
func StopThePool() {
	tasks.m.Lock()
	defer tasks.m.Unlock()

	tasks.isOk = false

}

func ActiveThePool() {
	tasks.m.Lock()
	defer tasks.m.Unlock()

	tasks.isOk = true
}

func GetTask(name string) (t *Task, err error) {
	if name != "" {
		t := tasks.pool[name]

		if t != nil {
			return t, nil
		} else {
			return nil, errors.New(fmt.Sprintf("the task %s is not in the pool ", name))
		}
	} else {
		return nil, errors.New("the name cannot be empty")
	}
}

func (t *Task) Pause() {

	go func() { t.Chan <- Pause }()

}

func (t *Task) Stop() {

	go func() { t.Chan <- Stop }()

}

func (t *Task) Start() {

	go t.Run()
	t.Chan <- Running
}

func Add2Pool(t *Task) error {
	tasks.m.Lock()
	defer tasks.m.Unlock()
	name := t.Name

	if tasks.pool[name] != nil {
		return errors.New(fmt.Sprintf("the task namede %s is exist", name))
	}

	tasks.pool[name] = t

	return nil
}

func DelFromPool(name string) error {

	if tasks.pool[name] == nil {
		return errors.New(fmt.Sprintf("no one task named %s in the pool", name))
	}

	delete(tasks.pool, name)
	return nil
}

func NewTask(name string, content func() error, duration time.Duration) (*Task, error) {
	if name != "" {
		t := &Task{
			Name:      name,
			Status:    Running,
			Duration:  duration,
			Content:   content,
			CreatTime: time.Now(),
			Chan:      make(chan string),
		}

		return t, nil
	} else {
		//isstart = false
		return nil, errors.New("the name can not be empty")
	}
}

//run the task single
func (t *Task) Run() {
	timer := time.NewTicker(t.Duration)
	for {
		select {
		case <-timer.C:
			if t.Status == Pause {
				runtime.Gosched()
				continue
			}

			if err := t.Content(); err != nil {
				timer.Stop()
				log.Println("task ", t.Name, "execute fail : ", err)
				return
			}

		case status, ok := <-t.Chan:
			if !ok {
				t.Status = Running
			}

			switch status {
			case Stop:
				timer.Stop()
				log.Println("stop task", t.Name)
				return
			case Pause:
				t.Status = Pause
				log.Println("pause task", t.Name)
			case Running:
				t.Status = Running
			}

		}
	}
}

func RunPool() error {
	if tasks.isOk {

		if tasks.pool == nil || len(tasks.pool) == 0 {
			return errors.New("no task in the pool")
		}

		for name, t := range tasks.pool {
			log.Println(fmt.Sprintf("task %s start!", name))
			go t.Run()
		}

		return nil
	} else {
		return errors.New("the taskPool is not ok!")
	}
}

/**
func main() {

	t, err := NewTask("test1", func() error {
		log.Println("test1")
		return nil
	}, 1*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	t2, err := NewTask("test2", func() error {
		log.Println("test2")
		return nil
	}, 3*time.Second)

	if err != nil {
		log.Fatal(err)
	}

	t3, err := NewTask("test3", func() error {
		log.Println("test3")
		return errors.New("yeah ,error")
	}, 3*time.Second)

	if err != nil {
		log.Fatal(err)
	}

	err = AddToPool("test1", t)
	if err != nil {
		log.Fatal(err)
	}
	err = AddToPool("test2", t2)
	if err != nil {
		log.Fatal(err)
	}

	err = AddToPool("test3", t3)
	if err != nil {
		log.Fatal(err)
	}

	t4, err := NewTask("test4", func() error {
		log.Println("test4")
		return nil
	}, 3*time.Second)

	if err != nil {
		log.Fatal(err)
	}

	t5, err := NewTask("test5", func() error {
		log.Println("test5")
		return nil
	}, 3*time.Second)

	if err != nil {
		log.Fatal(err)
	}

	t4.Start()

	err = RunPool()

	tt := time.NewTimer(10 * time.Second)

	select {
	case <-tt.C:
		t5.Stop()
		t4.Stop()

	}
	if err != nil {
		log.Fatal(err)
	}

	for {
	}
}
*/
