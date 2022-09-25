// Create data to database
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
	// Create Usergroup
	cdug, err := client.Debug().UserGroup.
		Create().
		SetName("group-a").
		Save(ct)
	if err != nil {
		log.Fatalf("failed create usergroup: %v", err)
	}
	log.Printf("UserGroup: %+v", cdug)
	// Create User
	cdu, err := client.Debug().User.
		Create().
		SetFirstName("Jason").
		SetLastName("Jones").
		SetMail("jason-jones@exp.comcom").
		SetUsergroupID(1).
		Save(ct)
	if err != nil {
		log.Fatalf("failed create user: %v", err)
	}
	log.Printf("User: %+v", cdu)

	log.Print("process done.")
}
