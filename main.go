package main

import (
	"context"
	"github.com/CristhoperDev/go-graphql-entc/internal/database"
)

func main()  {
	client, err := database.Open()
	if err != nil {
		panic(err)
	}
	err = client.Schema.Create(context.Background())
	if err != nil {
		panic(err)
	}
}
