package contact

import (
	"github.com/404th/Tasks/struct1/helper"
)

type Contact struct {
	ID          string
	Name        string
	PhoneNumber uint64
	Address     string
	Age         int
	Job         string
}

type ContactManager struct {
	Contacts []Contact
}

func (cm *ContactManager) Create(contact *Contact) Contact {
	contact.ID = helper.NewID()

	cm.Contacts = append(cm.Contacts, *contact)
	return *contact
}

func (cm ContactManager) Get(id string) Contact {
	var searchedContact Contact
	for i := 0; i < len(cm.Contacts); i++ {
		if cm.Contacts[i].ID == id {
			searchedContact = cm.Contacts[i]
			break
		}
	}
	return searchedContact
}

func (cm *ContactManager) Update(usr Contact) Contact {
	var updatedContact Contact
	for i := 0; i < len(cm.Contacts); i++ {
		if cm.Contacts[i].ID == usr.ID {
			cm.Contacts[i].Address = usr.Address
			cm.Contacts[i].Name = usr.Name
			cm.Contacts[i].PhoneNumber = usr.PhoneNumber
			cm.Contacts[i].Age = usr.Age
			cm.Contacts[i].Job = usr.Job
			updatedContact = cm.Contacts[i]
			break
		}
	}
	return updatedContact
}

func (cm *ContactManager) Delete(id string) Contact {
	var deletedContact Contact
	for i := 0; i < len(cm.Contacts); i++ {
		if id == cm.Contacts[i].ID {
			deletedContact = cm.Contacts[i]
			cm.Contacts = append(cm.Contacts[:i], cm.Contacts[i+1:]...)
			break
		}
	}
	return deletedContact
}

func (cts ContactManager) GetAll() []Contact {
	return cts.Contacts
}
