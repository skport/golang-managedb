// Testing connect to database
package main

import (
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
	log.Print("process done.")
}
