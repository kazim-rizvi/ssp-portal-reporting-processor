// csv/csv.go

package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"ssp-portal-reporting-processor/model"
)

type CSVWriter struct {
	filePath string
}

func NewCSVWriter(filePath string) *CSVWriter {
	return &CSVWriter{filePath: filePath}
}

func (cw *CSVWriter) WriteCSV(data []model.CreativeReviewModelOnDemand) error {
	file, err := os.Create(cw.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, item := range data {
		record := []string{fmt.Sprintf("%d", item.Id), item.Crid}
		err := writer.Write(record)
		if err != nil {
			return err
		}
	}

	return nil
}
