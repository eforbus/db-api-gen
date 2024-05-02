package main

import (
	"log"
	"net/http"

	"employees/ent"
	"employees/ent/ogent"

	"entgo.io/ent/dialect"

	//_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Create ent client.
	//client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	//if err != nil {
	//    log.Fatal(err)
	//}
	// Run the migrations.
	//if err := client.Schema.Create(context.Background(), schema.WithAtlas(true)); err != nil {
	//    log.Fatal(err)
	//}

	client, err := ent.Open(dialect.MySQL, "root:root@tcp(localhost:3306)/employees?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	// Start listening.
	srv, err := ogent.NewServer(ogent.NewOgentHandler(client))
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":3000", srv); err != nil {
		log.Fatal(err)
	}
}
