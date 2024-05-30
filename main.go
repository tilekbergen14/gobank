package main

import "log"

func main() {
	s := NewAPIServer(":3000")
	log.Println(s.Run())
}
