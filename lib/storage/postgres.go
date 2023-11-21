package storage

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

type postgresStore struct {
	Conn  *pgx.Conn
	batch *pgx.Batch
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
	instance.batch = &pgx.Batch{}
	return instance
}

func (pg *postgresStore) Close() {
	if pg.batch.Len() > 0 {
		results := pg.Conn.SendBatch(context.Background(), pg.batch)
		results.Close()
	}
	if pg.Conn != nil {
		pg.Conn.Close(context.Background())
	}
}

func (pg *postgresStore) initTables() {
	_, err := pg.Conn.Exec(context.TODO(), `CREATE TABLE IF NOT EXISTS public.cnae (
		id varchar NOT NULL,
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

func (pg *postgresStore) Write(i interface{}) error {
	t, ok := i.(Cnae)
	if ok {
		pg.enqueueCnae(t)
	}
	return nil
}

func (pg *postgresStore) enqueueCnae(c Cnae) {
	fmt.Println(c)
	pg.Conn.Prepare(context.Background(), "insertCnae", "INSERT INTO public.cnae (id, label) values ($1, $2);")
	pg.batch.Queue("insertCnae", c.Code, c.Label)
}
