package main

import (
	"log"
	"solution/server"
	"solution/service"
	"solution/store/postgresql"
)

func main() {
	db, err := postgresql.New()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	service := service.NewService(db)
	service.SetLastVersion("managed-k8s")

	s := server.NewServer(*service)

	err = s.Start()
	if err != nil {
		log.Fatalf("server is not running:%s", err)
	}
}
