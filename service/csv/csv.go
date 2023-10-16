// csv/csv.go

package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"ssp-portal-reporting-processor/model"

	"github.com/google/uuid"
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

func CreateCSVFile(headers []string) (string, error) {
	uuid := uuid.New()

	// Create the CSV file with the UUID in the file name.
	fileName := fmt.Sprintf("%s.csv", uuid)
	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}

	// Write the UTF-8 BOM to the file.
	bom := []byte{0xEF, 0xBB, 0xBF}
	file.Write(bom)

	// Create a CSV writer for the file.
	csvWriter := csv.NewWriter(file)

	// Write the headers to the CSV file.
	if err := csvWriter.Write(headers); err != nil {
		file.Close() // Close the file if there's an error
		return "", err
	}

	// Flush and close the CSV writer.
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		file.Close() // Close the file if there's an error
		return "", err
	}

	file.Close() // Close the file

	return fileName, nil
}

func WriteCSVRowsToCSVFile(filePath string, rows []string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	for _, row := range rows {
		err := writer.Write([]string{row})
		if err != nil {
			return err
		}
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		return err
	}

	return nil
}
