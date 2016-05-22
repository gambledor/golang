// Package main provides ...
package main

import (
	"database/sql"
	"github.com/gambledor/go-in-5-minutes/episode13/models"
	_ "github.com/mxk/go-sqlite/sqlite3"
	"log"
)

const (
	sqlite3Str = "sqlite3"
	memStr     = ":memory:"
)

func main() {
	// Change this line to change the database.
	db, err := sql.Open(sqlite3Str, memStr)
	if err != nil {
		log.Fatalf("error opening DB (%s)", err)
	}
	log.Printf("Creating new table")
	if _, err := models.CreatePersonTable(db); err != nil {
		log.Fatalf("Error creating table (%s)", err)
	}
	log.Printf("Created")

	me := models.Person{Firstname: "giuseppe", Lastname: "lo brutto", Age: 37}
	log.Printf("Inserting %+v into DB", me)
	if _, err := models.InsertPerson(db, me); err != nil {
		log.Fatalf("error inserting new person  into db (%s)", err)
	}
	log.Printf("Inserted")

	log.Printf("Selecting person from DB")
	selectMe := models.Person{}
	if err := models.SelectPerson(db, me.Firstname, me.Lastname, me.Age, &selectMe); err != nil {
		log.Fatalf("Error selecting person from DB (%s)", err)
	}
	log.Printf("Selected %+v from the DB", selectMe)

	log.Printf("Updating person into DB")
	updateMe := models.Person{
		Firstname: "Aaron",
		Lastname:  "Schlesinger",
		Age:       31,
	}
	if err := models.UpdatePerson(db, selectMe.Firstname, selectMe.Lastname, selectMe.Age, updateMe); err != nil {
		log.Fatalf("Error updating person int the DB (%s)", err)
	}

	log.Printf("Deleting person from DB")
	if delErr := models.DeletePerson(db, selectMe.Firstname, selectMe.Lastname, selectMe.Age); delErr != nil {
		log.Fatalf("Error deleting  person from DB (%s)", delErr)
	}
	log.Printf("deleted")
}
