package main

import (
	"fmt"
	"github.com/olegvel/life-sheet/pkg/handlers"
	"log"
)

func main() {
	fmt.Println("Life Sheet")
	fmt.Println("Structure and save the progress of your life!")

	s := handlers.NewServer(":12965")
	fmt.Println("Listening on :12965...")
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
