package main

import (
	"testing"

	hp "github.com/404th/struct/helper"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	user := Contact{
		ID:          hp.NewID(),
		Name:        hp.NewName(),
		PhoneNumber: hp.NewNumber(),
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

	err := user.GetAboutUser()
	assert.NoError(t, err)

	err = task.GetAboutTask()
	assert.NoError(t, err)
}
