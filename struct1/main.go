package main

import (
	"fmt"

	"github.com/404th/Tasks/struct1/contact"
	"github.com/404th/Tasks/struct1/helper"
)

func main() {
	var contactsList contact.ContactManager

	// CREATE
	contactData := contact.Contact{
		Name:        helper.NewName(),
		Address:     helper.NewAddress(),
		Job:         helper.NewJob(),
		Age:         helper.NewAge(),
		PhoneNumber: helper.NewNumber(),
	}
	contactsList.Create(&contactData)

	fmt.Println("================ CREATE ==================")
	for inx, i := range contactsList.Contacts {
		fmt.Printf("%v: %v\n", inx, i)
	}

	// GET
	g := contactsList.Get(contactData.ID)
	fmt.Println("================ GET ==================")
	fmt.Printf("GET => %v\n", g)

	// UPDATE
	newContact := contact.Contact{
		ID:      contactsList.Contacts[0].ID,
		Name:    "updated NAME",
		Address: "updated ADDRESS",
		Age:     1000,
		Job:     "updated JOB",
	}

	fmt.Println("================ UPDATE ==================")
	contactsList.Update(newContact)
	for inx, i := range contactsList.Contacts {
		if inx == 0 {
			fmt.Printf("===============> %v: %v \n", inx, i)
		} else {
			fmt.Printf("%v: %v\n", inx, i)
		}
	}

	// GET ALL
	fmt.Println("================ GET ALL ==================")
	data := contactsList.GetAll()
	for inx, i := range data {
		fmt.Printf("%v: %v\n", inx, i)
	}

	// DELETE
	deletedContact := contactsList.Delete(data[0].ID)
	fmt.Println("================ DELETE ==================")
	for inx, i := range contactsList.Contacts {
		fmt.Printf("%v: %v\n", inx, i)
	}
	fmt.Println("================ DELETED ITEM ==================")
	fmt.Printf("Deleted: %v\n", deletedContact)
}
