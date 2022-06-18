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

	faunatable := `CREATE TABLE fauna (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "name" TEXT,
        "scientificname" TEXT,
		"placesfound" TEXT,
        "description" TEXT);`
	query, err := DB.Prepare(faunatable)
	if err != nil {
		log.Fatal(err)
	}
	query.Exec()
	return nil
}

func AnimalsGot() ([]Animal, error) {

	rows, err := DB.Query("SELECT id, name, scientificname, placesfound, description from fauna")

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
