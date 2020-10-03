package main

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"invoices/pkg/database"
	"invoices/pkg/handler"
)

const (
	dbHost     = "localhost"
	dbPort     = "5436"
	dbUsername = "postgres"
	dbPassword = "qwerty"
	dbName     = "postgres"
	sslMode    = "disable"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	db, err := database.NewPostgresDB(database.Config{
		Host:     dbHost,
		Port:     dbPort,
		Username: dbUsername,
		Password: dbPassword,
		DBName:   dbName,
		SSLMode:  sslMode,
	})
	if err != nil{
		logrus.Fatalf("Error connecting to Database: %s", err.Error())
	}

	defer db.Close()

	repository := database.NewInvoiceRepository(db)

	h := handler.NewHandler(repository)
	if err := h.Init(); err != nil {
		logrus.Printf("Error occurred while running HTTP server: %s", err.Error())
	}
}
