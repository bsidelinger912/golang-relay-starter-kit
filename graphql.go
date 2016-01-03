package main

import (
	"log"
	"net/http"

	"github.com/bsidelinger912/golang-relay-starter-kit/models"
	"github.com/graphql-go/handler"
)

func main() {
	//connect the db here, this way the models are independant of their data source
	dbErr := models.NewDB("user=benjaminsidelinger dbname=mystery_shopper password=Cheet@h912 host=localhost port=5432 sslmode=disable")
	if dbErr != nil {
		log.Panic(dbErr)
	}

	// simplest relay-compliant graphql server HTTP handler
	h := handler.New(&handler.Config{
		Schema: &models.Schema,
		Pretty: true,
	})

	// create graphql endpoint
	http.Handle("/graphql", h)

	// serve!
	port := ":8080"
	log.Printf(`GraphQL server starting up on http://localhost%v`, port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("ListenAndServe failed, %v", err)
	}
}
