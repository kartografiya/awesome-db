package main

import (
	"log"

	"github.com/kartografiya/awesome-db/db"
)

func main() {
	v1 := db.Video{
		Title: "Inception",
		Director: "Chirstopher Nolan",
		Year: 2010,
	}
	v2 := db.Video{
		Title: "The Terminator",
		Director: "James Cameron",
		Year: 1984,
	}

	myDb := db.Init()
	myDb.Add(v1)
	myDb.Add(v2)
	log.Printf("My db size is %d\n", myDb.Size())

	log.Print(myDb.String())
}
