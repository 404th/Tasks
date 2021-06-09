package repository

import (
	"log"

	"github.com/404th/testit/genproto"
	"github.com/jmoiron/sqlx"
)

type ContactPg struct {
	db *sqlx.DB
}

func NewContactRepository(db *sqlx.DB) *ContactPg {
	return &ContactPg{
		db: db,
	}
}

func (r *ContactPg) CreateContact(ctc *genproto.ContactCreateReq) (genproto.WithID, error) {
	var (
		contact_id int32
		withId     genproto.WithID
	)
	query := `INSERT INTO contacts (name, username, age) VALUES ($1, $2, $3) RETURNING id`
	row := r.db.QueryRow(query, ctc.Name, ctc.Username, ctc.Age)
	if err := row.Scan(&contact_id); err != nil {
		log.Fatalf("error while creating contact: %v", err)
		return genproto.WithID{}, err
	}
	withId.Id = contact_id
	return withId, nil
}

func (r *ContactPg) GetContact(ctc *genproto.WithID) (genproto.Contact, error) {
	var contact genproto.Contact

	query := `SELECT name, username, age FROM contacts WHERE id=$1`
	row := r.db.QueryRow(query, ctc.Id)
	if err := row.Scan(&contact.Name, &contact.Username, &contact.Age); err != nil {
		log.Fatalf("error while getting contact: %v", err)
		return genproto.Contact{}, err
	}
	return contact, nil
}

func (r *ContactPg) GetAllContacts(ctc *genproto.EmptyResponse) (genproto.Contacts, error) {
	var contacts genproto.Contacts

	query := `SELECT * FROM contacts`
	rows, err := r.db.Query(query)
	if err != nil {
		return genproto.Contacts{}, err
	}

	for rows.Next() {
		var contact genproto.Contact
		rows.Scan(
			&contact.Id,
			&contact.Name,
			&contact.Username,
			&contact.Age,
		)

		contacts.Contacts = append(contacts.Contacts, &contact)
	}

	return contacts, nil
}
func (r *ContactPg) UpdataContact(ctc *genproto.Contact) (genproto.Contact, error) {
	var (
		contact genproto.Contact
		idd     int32
	)

	query := `UPDATE contacts SET name = $1, username = $2, age = $3 WHERE id = $4 RETURNING id;`

	row := r.db.QueryRow(query, ctc.Name, ctc.Username, ctc.Age, ctc.Id)
	if err := row.Scan(&idd); err != nil {
		log.Fatalf("error while updating contact: %v", err)
		return genproto.Contact{}, err
	}

	query1 := `SELECT name, username, age FROM contacts WHERE id=$1`
	row1 := r.db.QueryRow(query1, idd)
	if err := row1.Scan(&contact.Name, &contact.Username, &contact.Age); err != nil {
		log.Fatalf("error while getting contact: %v", err)
		return genproto.Contact{}, err
	}

	contact.Id = idd
	return contact, nil
}

func (r *ContactPg) DeleteContact(ctc *genproto.WithID) (genproto.EmptyResponse, error) {
	var idd string

	query := `DELETE FROM contacts WHERE id = $1 RETURNING id`

	row := r.db.QueryRow(query, ctc.Id)
	if err := row.Scan(&idd); err != nil {
		log.Fatalf("error while deleting contact: %v", err)
		return genproto.EmptyResponse{}, err
	}

	return genproto.EmptyResponse{}, nil
}
