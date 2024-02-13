package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

type Config struct {
	Mailer Mail
}

const webPort = ":80"

func main() {
	app := Config{
		Mailer: createMail(),
	}

	log.Println("Starting service on port", webPort)

	srv := &http.Server{
		Addr:    webPort,
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func createMail() Mail {
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	m := Mail{
		Domain:     os.Getenv("DOMAIN"),
		Host:       os.Getenv("MAIL_HOST"),
		Port:       port,
		Encryption: os.Getenv("MAIL_ENCRYPTION"),
		Username:   os.Getenv("MAIL_USERNAME"),
		Password:   os.Getenv("MAIL_PASSWORD"),
		FromName:   os.Getenv("FROM_NAME"),
		FromEmail:  os.Getenv("FROM_ADDRESS"),
	}

	return m
}
