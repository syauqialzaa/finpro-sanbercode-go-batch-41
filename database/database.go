package database

import (
	"database/sql"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	migrate "github.com/rubenv/sql-migrate"
)

func Migrations(db *sql.DB) {
	migrations := &migrate.PackrMigrationSource {
		Box: packr.New("migrations", "./sql"),
	}

	_, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		fmt.Println("[DEBUG] failed to execute sql file.")
		panic(err)
	}
}