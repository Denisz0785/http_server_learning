package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectDB(urlDb string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv(urlDb))
	if err != nil {
		err = fmt.Errorf("error to connetct DB %v", err)
		return nil, err
	}
	return conn, nil
}

// GetManyRows gets all rows of type of expenses from DB by login
func GetManyRowsByLogin(connec *pgx.Conn, login string) ([]string, error) {
	rows, _ := connec.Query(context.Background(), "SELECT surname from users")
	numbers, err := pgx.CollectRows(rows, pgx.RowTo[string])
	if err != nil {
		err = fmt.Errorf("ups: %v", err)
		return nil, err
	}
	return numbers, nil
}
