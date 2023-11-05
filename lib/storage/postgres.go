package storage

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

type postgresStore struct {
	Conn *pgx.Conn
}

func NewPostgresStore() *postgresStore {
	conn, err := connect()
	if err != nil {
		log.Fatalf("Failed to establish a connection with Postgres: %s\n", err.Error())
	}

	instance := &postgresStore{
		Conn: conn,
	}
	instance.initTables()
	return instance
}

func (pg *postgresStore) Close() {
	if pg.Conn != nil {
		pg.Conn.Close(context.Background())
	}
}

func (pg *postgresStore) initTables() {
	_, err := pg.Conn.Exec(context.TODO(), `CREATE TABLE IF NOT EXISTS public.cnae (
		id _varchar NOT NULL,
		"label" varchar NOT NULL,
		CONSTRAINT cnae_pk PRIMARY KEY (id)
	);`)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func (pg *postgresStore) GetCnaes() pgx.Rows {
	rows, err := pg.Conn.Query(context.Background(), "SELECT * from public.cnae limit 10;")
	if err != nil {
		fmt.Println(err.Error())
	}

	return rows
}

func connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn, err
}
