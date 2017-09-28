package pg

import (
	"database/sql"

	"fmt"
	_ "github.com/lib/pq"
)

func Run(cmd, conn string) error {
	db, err := sql.Open("postgres", conn)
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
