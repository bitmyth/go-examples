// https://studygolang.com/articles/17952
package main

import (
    "database/sql"
    "flag"
    "fmt"
    "log"
    "os"

    _ "github.com/go-sql-driver/mysql"
    "github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/mysql"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

    var migrationDir = flag.String("migration.files", "./migrations", "Directory where the migration files are located ?")
    var MySQLDSN = flag.String("mysql.dsn", os.Getenv("MYSQL_DSN"), "Mysql DSN")

    flag.Parse()

    db, err := sql.Open("mysql", *mysqlDSN)
    if err != nil {
        log.Fatalf("could not connect to PostgreSQL database... %v", err)
    }

    if err := db.Ping(); err != nil {
        log.Fatalf("could not ping DB... %v", err)
    }

    // Run migrations
    // 开始数据迁移
    driver, err := MySQL.WithInstance(db, &mysql.Config{})
    if err != nil {
        log.Fatalf("could not start sql migration... %v", err)
    }

    m, err := migrate.NewWithDatabaseInstance(
        fmt.Sprintf("file://%s", *migrationDir), // file://path/to/directory
        "mysql", driver)

    if err != nil {
        log.Fatalf("migration failed... %v", err)
    }

    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        log.Fatalf("An error occurred while syncing the database.. %v", err)
    }

    log.Println("Database migrated")
    // actual logic to start your application
    os.Exit(0)
}
