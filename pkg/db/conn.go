package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"medDods/MedodsGo/pkg/envRead"
)

var ctx = context.Background()

func connDB(config *envRead.AppConfig) (*pgxpool.Pool, error) {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)
	// postgres://user_name:password@localhost:5432/dbname
	conn, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return conn, nil
}

func initDB(pgx *pgxpool.Pool) {
	dropTable := "DROP TABLE IF EXISTS ;"
	rows, err := pgx.Query(ctx, dropTable)
	if err != nil {
		log.Fatal(err)
	}
	_ = rows
}
