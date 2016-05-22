// Package models provides the application models.
package models

import (
	"database/sql"
	"fmt"
)

const (
	PersonTableName    = "person"
	PersonFirstNameCol = "firstname"
	PersonLastnameCol  = "lastname"
	PersonAgeCol       = "age"
)

type Person struct {
	Firstname string
	Lastname  string
	Age       uint
}

// CreatePersonTable uses db to create a new person table for person model and returns the result.
func CreatePersonTable(db *sql.DB) (sql.Result, error) {
	return db.Exec(
		fmt.Sprintf(
			"CREATE TABLE %s (%s varchar(255), %s varchar(255), %s unsigned int)",
			PersonTableName,
			PersonFirstNameCol,
			PersonLastnameCol,
			PersonAgeCol,
		),
	)
}

// InserPerson insert person into a DB
func InsertPerson(db *sql.DB, person Person) (sql.Result, error) {
	return db.Exec(
		fmt.Sprintf("INSERT INTO %s VALUES (?, ?, ?)", PersonTableName),
		person.Firstname,
		person.Lastname,
		person.Age,
	)
}

// SelectPerson select a person with a ginve firstname, lastname and age. On success writes the result into result and
// on failure, return a non-nil error and makes no modification to result.
func SelectPerson(db *sql.DB, firstname, lastname string, age uint, result *Person) error {
	row := db.QueryRow(
		fmt.Sprintf(
			"SELECT * FROM %s WHERE %s = ? AND %s = ? AND %s = ?",
			PersonTableName,
			PersonFirstNameCol,
			PersonLastnameCol,
			PersonAgeCol,
		),
		firstname,
		lastname,
		age,
	)
	var retFistname, retLastname string
	var retAge uint
	if err := row.Scan(&retFistname, &retLastname, &retAge); err != nil {
		return err
	}
	result.Firstname = retFistname
	result.Lastname = retLastname
	result.Age = retAge

	return nil
}

// UpdatePerson updates the person with the given firs and last name and age with newPerson.
// Return a non-nil error if the update failed and nil if successeds.
func UpdatePerson(db *sql.DB, firstname, lastname string, age uint, newPerson Person) error {

	_, err := db.Exec(
		fmt.Sprintf(
			"UPDATE %s SET %s=?, %s=?, %s=? WHERE %s=? AND %s=? AND %s=?",
			PersonTableName,
			PersonFirstNameCol,
			PersonLastnameCol,
			PersonAgeCol,
			PersonFirstNameCol,
			PersonLastnameCol,
			PersonAgeCol,
		),
		newPerson.Firstname,
		newPerson.Lastname,
		newPerson.Age,
		firstname,
		lastname,
		age,
	)

	return err
}

// DeletePerson deletes the person with the given first and last name and age.
// Returns a non-nil error if the delete failed, and nil if the delete succeded.
func DeletePerson(db *sql.DB, firstname, lastname string, age uint) error {
	_, err := db.Exec(
		fmt.Sprintf(
			"DELETE FROM %s WHERE %s=? AND %s=? AND %s=?",
			PersonTableName,
			PersonFirstNameCol,
			PersonLastnameCol,
			PersonAgeCol,
		),
		firstname,
		lastname,
		age,
	)

	return err
}
