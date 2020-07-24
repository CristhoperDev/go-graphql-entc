package database

import (
	"github.com/CristhoperDev/go-graphql-entc/internal/provider/entc"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
)

func Open() (*entc.Client, error) {
	client, err := entc.Open(viper.GetString("DATABASE_DRIVER"), viper.GetString("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	//defer client.Close()

	return client, nil
}