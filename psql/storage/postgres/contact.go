package postgres

import (
	"github.com/404th/Tasks/psql/storage/models"
	"github.com/404th/Tasks/psql/storage/repo"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type contactRepo struct {
	db *sqlx.DB
}

func NewContactRepo(db *sqlx.DB) repo.ContactStorageI {
	return &contactRepo{db: db}
}

func (cc *contactRepo) Create(m models.Contact) (string, error) {
	id := uuid.New().String()

	query := `INSERT INTO contacts (
			id,
			name,
			username,
			phone_number,
			profession,
			age,
			is_admin ) VALUES ( $1, $2, $3, $4, $5, $6, $7 )
		`
	_, err := cc.db.Exec(query,
		id,
		m.Name,
		m.Username,
		m.PhoneNumber,
		m.Profession,
		m.Age,
		m.IsAdmin,
	)
	if err != nil {
		return "", nil
	}

	return id, err
}

func (cc *contactRepo) Get(id string) (models.Contact, error) {
	var contact models.Contact

	query := `SELECT id, name, username, phone_number, profession, age, is_admin FROM contacts WHERE id = $1`
	row := cc.db.QueryRow(query, id)
	err := row.Scan(&contact.ID, &contact.Name, &contact.Username, &contact.PhoneNumber, &contact.Profession, &contact.Age, &contact.IsAdmin)

	return contact, err
}

func (cc *contactRepo) GetAll() ([]models.Contact, error) {
	var contacts []models.Contact
	query := `SELECT * FROM contacts`
	err := cc.db.Select(&contacts, query)

	return contacts, err
}

func (cc *contactRepo) Update(m models.Contact) error {
	query := `UPDATE contacts
						SET name=$1,
						username=$2,
						phone_number=$3,
						profession=$4,
						age=$5,
						is_admin=$6 
						WHERE id=$7`
	_, err := cc.db.Exec(query,
		m.Name,
		m.Username,
		m.PhoneNumber,
		m.Profession,
		m.Age,
		m.IsAdmin,
		m.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (cc *contactRepo) Delete(id string) (string, error) {
	_, err := cc.db.Exec(`DELETE FROM contacts WHERE id=$1`, id)
	if err != nil {
		return "", err
	}

	return id, nil
}
