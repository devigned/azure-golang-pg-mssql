package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Run(cmd, conn string) error {
	db, err := sql.Open("mysql", conn)
	if err != nil {
		return err
	}

	rows, err := db.Exec(cmd)
	if err != nil {
		return err
	}

	fmt.Printf("Rows affected: %v", rows)
	return nil
}
