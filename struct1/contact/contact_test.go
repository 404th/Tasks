package contact_test

import (
	"fmt"
	"testing"

	"github.com/404th/Tasks/struct1/contact"
	"github.com/404th/Tasks/struct1/helper"
	"github.com/stretchr/testify/assert"
)

var testContactsList contact.ContactManager

func createUser(t *testing.T) contact.Contact {
	contactData := contact.Contact{
		Name:        helper.NewName(),
		Address:     helper.NewAddress(),
		Job:         helper.NewJob(),
		Age:         helper.NewAge(),
		PhoneNumber: helper.NewNumber(),
	}
	createdContact := testContactsList.Create(&contactData)
	assert.NotEmpty(t, createdContact)

	return createdContact
}

func deleteUser(t *testing.T, cr contact.Contact) contact.Contact {
	fmt.Println(testContactsList)
	deletedContact := testContactsList.Delete(cr.ID)
	fmt.Println(testContactsList)
	assert.NotEmpty(t, deletedContact, "error while deleting")
	return deletedContact
}

func TestCreate(t *testing.T) {
	createdContact := createUser(t)
	assert.NotEmpty(t, createdContact)
	deleteUser(t, createdContact)
}

func TestGet(t *testing.T) {
	createdContact := createUser(t)
	assert.NotEmpty(t, createdContact)

	data := testContactsList.Get(createdContact.ID)
	assert.NotEmpty(t, data, "error while getting")

	deleteUser(t, createdContact)
}

func TestGetAll(t *testing.T) {
	createdContact := createUser(t)
	assert.NotEmpty(t, createdContact)

	contacts := testContactsList.GetAll()
	assert.NotEmpty(t, contacts)

	deleteUser(t, createdContact)
}

func TestUpdate(t *testing.T) {
	createdContact := createUser(t)
	assert.NotEmpty(t, createdContact)

	newContact := contact.Contact{
		Name:        helper.NewName(),
		Address:     helper.NewAddress(),
		Job:         helper.NewJob(),
		Age:         helper.NewAge(),
		PhoneNumber: helper.NewNumber(),
	}
	newContact.ID = testContactsList.Contacts[0].ID
	updatedContact := testContactsList.Update(newContact)
	assert.NotEmpty(t, updatedContact)
	testContactsList.Delete(newContact.ID)
}
