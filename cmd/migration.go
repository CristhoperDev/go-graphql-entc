package cmd

import (
	"context"
	"github.com/CristhoperDev/go-graphql-entc/config"
	"github.com/CristhoperDev/go-graphql-entc/internal/database"
	"github.com/spf13/cobra"
)

func init()  {
	rootCmd.AddCommand(migrationCmd)
}

var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "run migrate database",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return config.ReadConfig()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		execute()
		return nil
	},
}

func execute()  {
	client, err := database.Open()
	if err != nil {
		panic(err)
	}
	err = client.Schema.Create(context.Background())
	if err != nil {
		panic(err)
	}
}