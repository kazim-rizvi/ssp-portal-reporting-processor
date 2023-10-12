// db/db.go

package db

import (
	"database/sql"
	"errors"
	"log"
	"ssp-portal-reporting-processor/config"
	"ssp-portal-reporting-processor/model"

	_ "github.com/go-sql-driver/mysql"
)

type DataFetcher struct {
	db *sql.DB
}

func NewDataFetcher(cfg config.DBConfig) (*DataFetcher, error) {
	// connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
	// 	cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	// db, err := sql.Open("mysql", connectionString)
	// if err != nil {
	// 	return nil, err
	// }

	// db.SetMaxOpenConns(1) // Set the maximum number of open connections
	// db.SetMaxIdleConns(1) // Set the maximum number of idle connections

	// return &DataFetcher{db: db}, nil
	return nil, nil
}

func (df *DataFetcher) FetchDataBatched(query string) ([]model.CreativeReviewModelLight, error) {
	if df.db == nil {
		return nil, errors.New("Database connection not initialized")
	}

	rows, err := df.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []model.CreativeReviewModelLight
	for rows.Next() {
		var d model.CreativeReviewModelLight
		err := rows.Scan(&d.Id, &d.Crid)
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
