package config

import (
	"context"
	"log"
	"os"

	model "back/models"
	util "back/utils"

	driver "github.com/arangodb/go-driver"

	"github.com/arangodb/go-driver/http"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBconnection struct {
	Connect driver.Database
}

func (db *DBconnection) CreateConnection() {

	var err error

	var conn driver.Connection

	databaseURI := make(chan string, 1)

	if os.Getenv("GO_ENV") != "production" {
		databaseURI <- util.GodotEnv("DATABASE_URI_DEV")
	} else {
		databaseURI <- os.Getenv("DATABASE_URI_PROD")
	}

	log.Printf("GO_ENV:%v", os.Getenv("GO_ENV"))
	log.Printf("URI_PROD:%v", os.Getenv("DATABASE_URI_PROD"))

	conn, err = http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{<-databaseURI},
		//Endpoints: []string{"https://5a812333269f.arangodb.cloud:8529/"},
	})
	if err != nil {
		log.Fatalf("Failed to create HTTP connection: %v", err)
	}
	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("root", "letmein"),
		//Authentication: driver.BasicAuthentication("root", "wnbGnPpCXHwbP"),
	})

	if err != nil {
		defer logrus.Info("Connection to Database Failed")
		logrus.Fatal(err.Error())
	}

	if os.Getenv("GO_ENV") != "production" {
		logrus.Info("Connection to Database Successfully")
	}
	ctx := context.Background()
	dbc, err := client.Database(ctx, "smacktalk")

	if err != nil {
		defer logrus.Info("Connection to smacktalk Failed")
		logrus.Fatal(err.Error())
	}

	db.Connect = dbc

}

func _Connection() *gorm.DB {
	databaseURI := make(chan string, 1)

	if os.Getenv("GO_ENV") != "production" {
		databaseURI <- util.GodotEnv("DATABASE_URI_DEV")
	} else {
		databaseURI <- os.Getenv("DATABASE_URI_PROD")
	}

	db, err := gorm.Open(postgres.Open(<-databaseURI), &gorm.Config{})

	if err != nil {
		defer logrus.Info("Connection to Database Failed")
		logrus.Fatal(err.Error())
	}

	if os.Getenv("GO_ENV") != "production" {
		logrus.Info("Connection to Database Successfully")
	}

	err = db.AutoMigrate(
		&model.EntityUsers{},
		&model.EntityStudent{},
	)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return db
}
