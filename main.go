package main

import (
	"log"
)

func main() {
	store, err := newPostgresStorage()
	if err != nil {
		log.Fatal(err)
	}
	if err := store.init(); err != nil {
		log.Fatal(err)
	}
	s := NewAPIServer(":3000", store)
	log.Println(s.Run())
}
