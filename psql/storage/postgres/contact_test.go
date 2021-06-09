package postgres_test

import (
	"math/rand"
	"testing"

	"github.com/404th/Tasks/psql/storage/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func createContact(t *testing.T) string {
	contact_id, err := strg.Contact().Create(
		models.Contact{
			Name:        uuid.NewString(),
			Username:    uuid.NewString(),
			Age:         rand.Intn(18 + (78 - 18 + 1)),
			PhoneNumber: uuid.NewString(),
			Profession:  uuid.NewString(),
			IsAdmin:     rand.Intn(2) == 1,
		},
	)
	assert.NoError(t, err)

	return contact_id
}

func deleteContact(t *testing.T, id string) {
	contact_id, err := strg.Contact().Delete(id)
	assert.NoError(t, err)
	assert.NotEmpty(t, contact_id)
}

func TestGet(t *testing.T) {
	contact_id := createContact(t)

	getContact, err := strg.Contact().Get(contact_id)
	assert.NoError(t, err)
	assert.NotEmpty(t, getContact)

	deleteContact(t, contact_id)
}
func TestGetAll(t *testing.T) {

	contact_id1 := createContact(t)
	contact_id2 := createContact(t)
	contact_id3 := createContact(t)

	contacts, err := strg.Contact().GetAll()
	assert.NoError(t, err)
	assert.NotEmpty(t, contacts)

	deleteContact(t, contact_id1)
	deleteContact(t, contact_id2)
	deleteContact(t, contact_id3)
}
func TestUpdate(t *testing.T) {

	contact_id := createContact(t)
	err := strg.Contact().Update(
		models.Contact{
			ID:          contact_id,
			Name:        uuid.NewString(),
			Username:    "Updated username" + uuid.NewString(),
			Age:         rand.Intn(80 + (100 - 80 + 1)),
			PhoneNumber: "updated phone number",
			Profession:  "Updated profession",
			IsAdmin:     false,
		},
	)
	assert.NoError(t, err)

	getContact, err := strg.Contact().Get(contact_id)
	assert.NoError(t, err, "error while getting contact in process updating")
	assert.NotEmpty(t, getContact)

	assert.Equal(t, getContact.PhoneNumber, "updated phone number")
	assert.Equal(t, getContact.Profession, "Updated profession")
	assert.Equal(t, getContact.IsAdmin, false)

	deleteContact(t, contact_id)
}
