package database

import (
	"github.com/CristhoperDev/go-graphql-entc/internal/provider/entc"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Open() (*entc.Client, error) {
	client, err := entc.Open("mysql", "root:root@tcp(localhost:3306)/ent_graphql?parseTime=true")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	//defer client.Close()

	return client, nil
}