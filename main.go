package main

import (
	"fmt"
	"log"
)

func main() {
	store, err := newPostgresStorage()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("%+v\n", store)
	// s := NewAPIServer(":3000", store)
	// log.Println(s.Run())
}
