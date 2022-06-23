package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Animal struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	ScientificName string `json:"scientificname"`
	Places         string `json:"placesfound"`
	Description    string `json:"description"`
}

var DB *sql.DB

func ConnectDb() error {
	db, err := sql.Open("sqlite3", "./database/wildlife.db")

	if err != nil {
		fmt.Println("Ok")
		return err
	}

	DB = db

	wildlifetable := `CREATE TABLE IF NOT EXISTS wildlife (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "name" TEXT,
        "scientificname" TEXT,
		"placesfound" TEXT,
        "description" TEXT);`
	query, err := DB.Prepare(wildlifetable)
	if err != nil {
		log.Fatal(err)
	}
	query.Exec()
	return nil
}

func AnimalsGot() ([]Animal, error) {

	rows, err := DB.Query("SELECT id, name, scientificname, placesfound, description from wildlife")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	Wildlife := make([]Animal, 0)

	for rows.Next() {
		singleAnimal := Animal{}
		err = rows.Scan(&singleAnimal.Id, &singleAnimal.Name, &singleAnimal.ScientificName, &singleAnimal.Places, &singleAnimal.Description)

		if err != nil {
			return nil, err
		}

		Wildlife = append(Wildlife, singleAnimal)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return Wildlife, err
}

func AnimalGotById(id string) (Animal, error) {
	stmt, err := DB.Prepare("SELECT id, name, scientificname, placesfound, description from wildlife WHERE id = ?")

	if err != nil {
		return Animal{}, err
	}

	animal := Animal{}

	sqlErr := stmt.QueryRow(id).Scan(&animal.Id, &animal.Name, &animal.ScientificName, &animal.Places, &animal.Description)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return Animal{}, nil
		}
		return Animal{}, sqlErr
	}

	return animal, err
}

func AnimalAdded(newAnimal Animal) (bool, error) {
	tx, err := DB.Begin()

	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("INSERT INTO wildlife (name, scientificname, placesfound, description) VALUES (?, ?, ?, ?)")

	defer stmt.Close()

	_, err = stmt.Exec(newAnimal.Name, newAnimal.ScientificName, newAnimal.Places, newAnimal.Description)

	tx.Commit()

	return true, nil
}

func AnimalUpdated(animalUpdate Animal, id int) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("ok")
		return false, err
	}

	stmt, err := tx.Prepare("UPDATE wildlife SET name = ?, scientificname = ?, placesfound = ?, description = ? WHERE id = ?")

	if err != nil {
		fmt.Println("ok2")
		fmt.Println(err)

		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(animalUpdate.Name, animalUpdate.ScientificName, animalUpdate.Places, animalUpdate.Description, id)

	if err != nil {
		fmt.Println("ok3")

		return false, err
	}

	tx.Commit()

	return true, nil
}

func AnimalDeleted(animalId int) (bool, error) {
	tx, err := DB.Begin()

	if err != nil {
		return false, err
	}

	stmt, err := DB.Prepare("DELETE from wildlife where id = ?")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(animalId)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}
