package main

import (
	"database/sql"
	"log"

	_ "gopkg.in/cq.v1"
)

func main() {
	db, err := sql.Open("neo4j-cypher", "http://neo4j:testings@localhost:7474")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare(`
        MATCH (d:Database)-[:SAYS]->(m:Message)
        WHERE d.name = {0}
        RETURN m.name
    `)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query("Neo4j")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var s string
	for rows.Next() {
		err := rows.Scan(&s)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(s)
	}
}
