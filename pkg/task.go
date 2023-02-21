package pkg

import (
	"fmt"

	"github.com/0xk2/todocli/third_party/utils"
)

const (
	TaskJustCreated string = "Task Just Created"
	TaskStarted     string = "Task Doing"
	TaskPaused      string = "Task Paused"
	TaskStopped     string = "Task Stopped"
	TaskFinished    string = "Task Finished"
)

type Task struct {
	Id          string
	Title       string
	Description string
	State       string
}

func (this *Task) ID() string {
	return this.Id
}

func NewTask(title string, description string) *Task {
	id := utils.RandomString(8)

	t := Task{
		Id:          id,
		Title:       title,
		Description: description,
		State:       TaskJustCreated,
	}
	fmt.Println(t)
	b := saveTask(&t)
	if b {
		return &t
	} else {
		return nil
	}
}

func (this *Task) Show() {
	fmt.Println("Task ID: ", this.Id)
	fmt.Println("---")
	fmt.Println(this.Title)
	fmt.Println(this.Description)
	fmt.Println("---")
	fmt.Println("Current State: ", this.State)
}

func (this *Task) ChangeStatus(code int8) {
	// 1 to start / doing
	// 0 to pause
	// 2 to finish
	// 3 to stop
	if this.State == TaskFinished || this.State == TaskStopped {
		fmt.Println(this.State, " cannot change anything")
		return
	} else if this.State == TaskJustCreated {
		if code == 1 {
			this.State = TaskStarted
		} else {
			fmt.Println(this.State, " you can only start it or leave it alone")
			return
		}
	} else if this.State == TaskStarted {
		if code == 0 {
			this.State = TaskPaused
		} else if code == 2 {
			this.State = TaskFinished
		} else if code == 3 {
			this.State = TaskStopped
		} else {
			fmt.Println(this.State, " already doing")
			return
		}
	} else if this.State == TaskPaused {
		if code == 3 {
			this.State = TaskStopped
		} else if code == 1 {
			this.State = TaskStarted
		} else if code == 2 {
			this.State = TaskFinished
		} else {
			fmt.Println(this.State, " already paused")
			return
		}
	}
	saveTask(this)
	this.Show()
}
