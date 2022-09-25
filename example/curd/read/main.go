// Update data to database
package main

import (
	"context"
	"fmt"
	"log"

	"managedb/example/config"

	"managedb/ent"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
			// user, pass, host, port, database
			config.DB["user"],
			config.DB["password"],
			config.DB["host"],
			config.DB["port"],
			config.DB["database"],
		),
	)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	// Execute
	ct := context.Background()
	// Select Users
	cdu, err := client.Debug().User.
		Query().
		All(ct)
	if err != nil {
		log.Fatalf("failed select users: %v", err)
	}
	log.Printf("Users: %+v", cdu)

	log.Print("process done.")
}
