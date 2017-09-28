package cmd

import (
	"errors"
	"github.com/devigned/azure-golang-pg-mssql/pg"
	"github.com/spf13/cobra"
)

var (
	pgCmd = &cobra.Command{
		Use:   "pg",
		Short: "Run a pg command",
		Long:  `Run a command against Azure PostgreSQL`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if command != "" || connString != "" {
				return nil
			} else {
				return errors.New("please provide a command and a connection string")
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return pg.Run(command, connString)
		},
	}
	command    string
	connString string
)

func init() {
	rootCmd.AddCommand(pgCmd)

	pgCmd.Flags().StringVarP(
		&command,
		"cmd",
		"c",
		"select 'hello world!'",
		"Command to execute against PostgreSQL")

	pgCmd.Flags().StringVarP(
		&connString,
		"conn",
		"n",
		"",
		"Connection string for PostgreSQL")
}
