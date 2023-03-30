package config

import (
	"context"
	"log"
	"os"
	"time"

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

	log.Printf("DB CONNECTION GO_ENV SETTING:%v:::DB URI:%v", os.Getenv("GO_ENV"), util.GodotEnv("DATABASE_URI_DEV"))

	var uptime = 10
	var client driver.Client

	conn, err = http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{<-databaseURI},
	})

	if err != nil {
		logrus.Fatal("Creation of connection string to failed", err.Error())
	}

	client, err = driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("root", "letmein"),
		//Authentication: driver.BasicAuthentication("root", "wnbGnPpCXHwbP"),
	})
	if err != nil {
		logrus.Fatal("Creation of NewClient failed", err.Error())
	}

	client.Connection().SetAuthentication(driver.BasicAuthentication("root", "letmein"))

	for i := 0; i < uptime; i++ {
		log.Printf("Attempting arangodb connection to smacktalk db")
		time.Sleep(5 * time.Second)

		ctx := context.Background()
		what, err := client.DatabaseExists(ctx, "smacktalk")
		if err != nil {
			logrus.Info("Try again. Smacktalk DB isnt found:   ", err.Error())

			if i == uptime {
				logrus.Fatal("Smacktalk DB isnt found after x times", err.Error())
			}
		} else {
			logrus.Info("Found smacktalk:", what)
			//worked
			i = 10
		}

	}

	ctx := context.Background()

	huh, err := client.Databases(ctx)
	if err != nil {
		log.Printf("ERROR:", err.Error())
	} else {

		log.Printf("huh:", huh)
	}

	users, err := client.Users(ctx)
	if err != nil {
		log.Printf("ERROR:", err.Error())
	} else {

		log.Printf("users:", users)
	}
	dbca, err := client.Databases(ctx)
	if err != nil {
		log.Printf("DBS ERROR:", err.Error())
	} else {
		log.Printf("DBS:", dbca)
	}

	dbc, err := client.Database(ctx, "smacktalk")

	if err != nil {
		defer logrus.Info("Connection to smacktalk Failed")
		logrus.Fatal("Smacktalk database not reachable", err.Error())
	} else {
		logrus.Info("Connection to Database Successfully")
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
	logrus.Info("Opening database: %s", databaseURI)
	db, err := gorm.Open(postgres.Open(<-databaseURI), &gorm.Config{})

	if err != nil {
		defer logrus.Info("Connection to Database Failed")
		logrus.Fatal(err.Error())
	}

	err = db.AutoMigrate(
		&model.EntityUsers{},
		&model.EntityStudent{},
	)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	if os.Getenv("GO_ENV") != "production" {
		logrus.Info("Connection to Database Successfully")
	}

	return db
}
