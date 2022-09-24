// Testing connect to database
package main

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/examples/start/ent"

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
			"mysql",
		),
	)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	log.Print("process done.")
}
