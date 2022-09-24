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
	// Delete User
	cdu, err := client.Debug().User.
		Delete().
		Where(user.FirstName("Jason")).
		Exec(ct)
	if err != nil {
		log.Fatalf("failed delete user: %v", err)
	}
	log.Printf("Deletetd : %+v", cdu)

	log.Print("process done.")
}
