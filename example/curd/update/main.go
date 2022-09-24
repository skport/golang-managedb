// Delete data to database
package main

import (
	"context"
	"fmt"
	"log"

	"managedb/ent"
	"managedb/ent/user"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
			// user, pass, host, port, database
			"root",
			"pass",
			"localhost",
			"3306",
			"testdb",
		),
	)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	// Execute
	ct := context.Background()
	// Update User
	cdu, err := client.Debug().User.
		Update().
		SetFirstName("Alex").
		Where(user.FirstName("Jason")).
		Save(ct)
	if err != nil {
		log.Fatalf("failed update user: %v", err)
	}
	log.Printf("Updated : %+v", cdu)

	log.Print("process done.")
}
