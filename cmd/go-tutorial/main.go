package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Esusu2017/rrs-commons/mongo"
	"github.com/mmuoDev/go-tutorial/internal/app"
)

func main() {
	port := "5050"
	os.Setenv("RRS_MONGO", "mongodb://127.0.0.1:27017")
	os.Setenv("RRS_DB_NAME", "test")
	provideMongoDB, err := mongo.NewConfigFromEnvVars().ToProvider(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	a := app.New(provideMongoDB)
	log.Println(fmt.Sprintf("Starting server on port:%s", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), a.Handler()))
}
