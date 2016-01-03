package test

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=benjaminsidelinger dbname=mystery_shopper password=Cheet@h912 host=localhost port=5432 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	var (
		id   int
		name string
	)
	rows, err := db.Query("select id, name from shoppers")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
