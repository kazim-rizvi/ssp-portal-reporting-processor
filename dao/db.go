package dao

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"ssp-portal-reporting-processor/model"

	_ "github.com/go-sql-driver/mysql"
)

type DataFetcher struct {
	db *sql.DB
}

func NewDataFetcher(connection model.DbConnection) (*DataFetcher, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		connection.User, connection.Password, connection.Host, connection.Port, connection.Database)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	return &DataFetcher{db: db}, nil
}

func FetchData[M any](df *DataFetcher, query string, addresses []*interface{}) ([]M, error) {
	if df.db == nil {
		return nil, errors.New("Database connection not initialized")
	}

	rows, err := df.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []M
	for rows.Next() {
		var d M
		err := rows.Scan(addresses)
		if err != nil {
			return nil, err
		}
		data = append(data, d)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func (df *DataFetcher) CloseConnection() error {
	if df.db != nil {
		log.Println("DB connection closed")
		return df.db.Close()
	}
	return nil
}
