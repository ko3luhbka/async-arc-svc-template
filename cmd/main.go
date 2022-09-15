package main

import (
	"log"

	"github.com/ko3luhbka/async-arc-svc-template/db"
	"github.com/ko3luhbka/async-arc-svc-template/migrations"
	"github.com/ko3luhbka/async-arc-svc-template/rest"
)

const (
	serviceName = "service1"
)

func main() {
	log.Printf("Starting service %s", serviceName)

	conn, err := db.NewConnection()
	if err != nil {
		log.Panic(err)
	}

	if err = db.RunMigrations(conn.DB, migrations.MigrationFiles); err != nil {
		log.Panic(err)
	}

	// TODO: kafka client init

	if err = rest.NewApp(); err != nil {
		log.Panic(err)
	}
}
