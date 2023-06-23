package app

import (
	"fmt"
	"issue-detector/pkg/handler"
	"issue-detector/pkg/repository"
	"log"
	"net/http"
)

func Init(config Config) error {

	hostsDb, err := repository.NewDatabase(config.HostsDB)
	if err != nil {
		log.Fatal(err)
	}
	usersDb, err := repository.NewDatabase(config.UsersDB)
	if err != nil {
		log.Fatal(err)
	}
	checkIpDb, err := repository.NewDatabase(config.CheckIpDB)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(hostsDb, usersDb, checkIpDb)
	handlers := handler.NewHandler(repo)
	server := http.Server{
		Addr:    ":" + config.Port,
		Handler: handlers,
	}

	fmt.Printf("Server started on port(%s) ...", config.Port)
	log.Fatal(server.ListenAndServe())
	return nil
}
