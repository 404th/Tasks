package main

import (
	"fmt"
	"log"

	"github.com/404th/Tasks/psql/storage"
	"github.com/404th/Tasks/psql/storage/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error while initializing configs: %s", err.Error())
	}

	// config := storage.Config{
	// 	Host:     viper.GetString("db.db_host"),
	// 	Port:     viper.GetString("db.db_port"),
	// 	Password: os.Getenv("DB_PASSWORD"),
	// 	Username: viper.GetString("db.username"),
	// 	DBName:   viper.GetString("db.db_name"),
	// 	SSLMode:  "disable",
	// }

	connStr := "user=spenc dbname=test_db password=mysecretpassword host=localhost sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nSuccessfully connected to database!\n")

	// db, err := storage.NewPostgresDB(config)

	// if err != nil {
	// 	log.Fatalf("unable to connect db: %s", err.Error())
	// }

	strg := storage.NewStoragePg(db)

	// CREATE
	contact_id, err := strg.Contact().Create(models.Contact{
		Name:        "Humoyun",
		Username:    uuid.NewString(),
		Age:         22,
		PhoneNumber: "998934642001",
		Profession:  "Teacher",
		IsAdmin:     false,
	})

	if err != nil {
		log.Fatalf("error while creating contact: %s", err.Error())
	}

	fmt.Printf("After CREATE contact: %v\n", contact_id)
	// GET
	getContact, err := strg.Contact().Get(contact_id)
	if err != nil {
		log.Fatalf("error while getting contact: %s", err.Error())
	}

	fmt.Printf("After getting => %v\n", getContact)

	// GET ALL
	contacts, err := strg.Contact().GetAll()
	if err != nil {
		log.Fatalf("error while getting all contacts: %s", err.Error())
	}

	fmt.Printf("After getting all => %v\n", contacts)

	// UPDATE
	err = strg.Contact().Update(
		models.Contact{
			ID:          getContact.ID,
			Name:        "Umarbek",
			Username:    "Umarbek",
			Age:         32,
			PhoneNumber: "998919191919",
			Profession:  "Professor",
			IsAdmin:     false,
		},
	)
	if err != nil {
		log.Fatalf("error while updating contacts: %s", err.Error())
	}

	// DELETE
	deleted_id, err := storage.NewStoragePg(db).Contact().Delete(getContact.ID)
	if err != nil {
		log.Fatalf("error while deleting contacts: %s", err.Error())
	}

	fmt.Printf("After deleted: deleted_id => %v\n", deleted_id)

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
