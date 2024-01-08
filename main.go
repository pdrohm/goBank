package main

import (
	"fmt"
	"log"
)


func main() {
	store, err := NewPostgressStore()
	if err != nil {
		log.Fatal(err)
	}


	fmt.Println("%+v\n", store)
}