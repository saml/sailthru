package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/saml/sailthru"
)

func main() {
	c := &sailthru.Client{
		Key:    os.Getenv("SAILTHRU_KEY"),
		Secret: os.Getenv("SAILTHRU_SECRET"),
	}

	c.Init("default account")
	b, err := json.Marshal(c.Lists)
	if err != nil {
		panic(err)
	}
	log.Printf("all lists: %v", string(b))

	email := os.Getenv("EMAIL")
	u, err := c.FetchUser(email)
	if err != nil {
		log.Fatalf("cannot fetch user: %v", email)
	}

	b, err = json.Marshal(u)
	if err != nil {
		panic(err)
	}
	log.Printf("fetched user: %v", string(b))

}
