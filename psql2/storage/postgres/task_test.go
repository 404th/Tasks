package postgres_test

import (
	"math/rand"
	"testing"

	"github.com/404th/psql2/storage/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func createTask(t *testing.T) string {
	id, err := strg.Task().Create(models.Task{
		ID:    "",
		Title: uuid.NewString(),
		Body:  uuid.NewString(),
		Done:  rand.Intn(2) == 1,
	})
	assert.NoError(t, err, "error while creating task")
	assert.NotEmpty(t, id, "after creating id should be returned")

	return id
}

func deleteTask(t *testing.T, id string) {
	delete_id, err := strg.Task().Delete(id)
	assert.NoError(t, err, "error while deleting task")
	assert.NotEmpty(t, delete_id, "after deleting item, id did not come")
}

func TestGet(t *testing.T) {
	id := createTask(t)
	task, err := strg.Task().Get(id)
	assert.NoError(t, err)
	assert.NotEmpty(t, task)
	deleteTask(t, id)
}

func TestGetAll(t *testing.T) {
	id1 := createTask(t)
	id2 := createTask(t)
	id3 := createTask(t)

	tasks, err := strg.Task().GetAll()
	assert.NoError(t, err)
	assert.NotEmpty(t, tasks)

	deleteTask(t, id1)
	deleteTask(t, id2)
	deleteTask(t, id3)
}

func TestUpdate(t *testing.T) {

	id := createTask(t)

	updatedTask := models.Task{
		ID:    id,
		Title: uuid.NewString(),
		Body:  "updated body",
		Done:  rand.Intn(2) == 1,
	}

	task, err := strg.Task().Update(updatedTask)
	assert.NoError(t, err)
	assert.NotEmpty(t, task)
	assert.EqualValues(t, task, id)

	deleteTask(t, id)
}
