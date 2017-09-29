package cmd

import (
	"errors"
	"github.com/devigned/azure-golang-pg-mssql/mysql"
	"github.com/spf13/cobra"
)

var (
	mysqlCmd = &cobra.Command{
		Use:   "mysql",
		Short: "Run a MySQL command",
		Long:  `Run a command against Azure MySQL`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if mysqlCmdString == "" || mysqlConnString == "" {
				return errors.New("please provide a command and a connection string")
			} else {
				return nil
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return mysql.Run(mysqlCmdString, mysqlConnString)
		},
	}

	mysqlCmdString  string
	mysqlConnString string
)

func init() {
	rootCmd.AddCommand(mysqlCmd)

	mysqlCmd.Flags().StringVarP(
		&mysqlCmdString,
		"cmd",
		"c",
		"select 'hello world!'",
		"Command to execute against MySQL")

	mysqlCmd.Flags().StringVarP(
		&mysqlConnString,
		"conn",
		"n",
		"",
		"Connection string for MySQL")
}
