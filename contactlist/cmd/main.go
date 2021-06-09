package main

import (
	"log"
	"os"

	"github.com/404th/contactlist"
	"github.com/404th/contactlist/pkg/handler"
	"github.com/404th/contactlist/pkg/repo"
	"github.com/404th/contactlist/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	// init config
	if err := initConfig(); err != nil {
		log.Fatalf("error occured while initializing config file: %s", err)
	}

	// init .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error occured while loading dotenv: %s", err)
	}

	db, err := repo.NewPostgresDB(
		repo.Config{
			Host:     viper.GetString("db.db_host"),
			Port:     viper.GetString("db.db_port"),
			Password: os.Getenv("DB_PASSWORD"),
			Username: viper.GetString("db.username"),
			DBName:   viper.GetString("db.db_name"),
			SSLMode:  viper.GetString("db.ssl_mode"),
		},
	)
	if err != nil {
		log.Fatalf("error while conntecting to db: %s", err)
	}

	repos := repo.NewRepo(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(contactlist.Server)

	if err := srv.Run(viper.GetString("1144"), handlers.InitRoutes()); err != nil {
		log.Fatalf("err occured while connecting to server: %s", err)
	}
}

// initialize Config file function
func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
