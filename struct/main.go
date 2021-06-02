package main

import (
	"errors"
	"fmt"
	"math/rand"

	hp "github.com/404th/struct/helper"
)

var Users []Contact
var Tasks []Task

type Contact struct {
	ID          string
	Name        string
	PhoneNumber uint64
	Address     string
	Age         int
	Job         string
}

type Task struct {
	ID        string
	ContactID string
	TaskList  []string
}

// know about user
func (f Contact) GetAboutUser() error {
	for r := 0; r < len(Users); r++ {
		if f.ID == Users[r].ID {
			fmt.Println("=========== About User ===========")
			fmt.Printf("Name: %v;\n Age: %v;\n Job: %v;\n Address: %v;\n", Users[r].Name, Users[r].Age, Users[r].Job, Users[r].Address)
			fmt.Println("==================================")
		}
	}
	return nil
}

// know about task
func (f Task) GetAboutTask() error {
	for r := 0; r < len(Tasks); r++ {
		if f.ID == Tasks[r].ID {
			fmt.Println("============ About task ==========")
			for idx, data := range Tasks[r].TaskList {
				fmt.Printf("%v : %v\n", idx, data)
			}
			fmt.Println("===================================")
		}
	}
	return nil
}

// get user by id
func GetUserById(id string) (Contact, error) {
	var usr Contact
	for r := 0; r < len(Users); r++ {
		if id == Users[r].ID {
			usr = Users[r]
		} else {
			return Contact{}, errors.New("user not found")
		}
	}
	return usr, nil
}

// get task by id
func GetTaskById(id string) (Task, error) {
	var tsk Task
	for r := 0; r < len(Tasks); r++ {
		if id == Tasks[r].ID {
			tsk = Tasks[r]
		} else {
			return Task{}, errors.New("task not found")
		}
	}
	return tsk, nil
}

func main() {
	// You can add below a number of tasks and users
	for a := 0; a < 100; a++ {
		user := Contact{
			ID:          hp.NewID(),
			Name:        hp.NewName(),
			PhoneNumber: hp.NewNumber(),
			Job:         hp.NewJob(),
			Address:     hp.NewAddress(),
			Age:         hp.NewAge(),
		}

		task := Task{
			ID:        hp.NewID(),
			ContactID: user.ID,
			TaskList:  hp.RandomTask(),
		}

		Users = append(Users, user)
		Tasks = append(Tasks, task)
	}

	us1 := Users[rand.Intn(len(Users))]
	ts1 := Tasks[rand.Intn(len(Tasks))]

	us1.GetAboutUser()
	ts1.GetAboutTask()
}
