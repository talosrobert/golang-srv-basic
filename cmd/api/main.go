package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const version = "0.0.1"

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "dbconn", os.Getenv("APPL_POSTGRES_DSN"), "PostgreSQL DSN")
	flag.Parse()

	db, err := openDB(cfg.db.dsn)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	appl := newApplication(cfg, db)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      appl.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
