package utils

import (
	"encoding/csv"
	"os"
)

// ExportToCSV writes a list of rows (each row = slice of strings) to a CSV file with a header
func ExportToCSV(headers []string, rows [][]string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write headers
	if err := writer.Write(headers); err != nil {
		return err
	}

	// Write rows
	return writer.WriteAll(rows)
}
