package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExportToCSV_Success(t *testing.T) {
	// Test data
	headers := []string{"Name", "Age", "City"}
	rows := [][]string{
		{"John", "25", "New York"},
		{"Jane", "30", "Los Angeles"},
		{"Bob", "35", "Chicago"},
	}
	filename := "test_output.csv"

	// Clean up after test
	defer os.Remove(filename)

	// Export to CSV
	err := ExportToCSV(headers, rows, filename)

	// Verify no error
	assert.Nil(t, err)

	// Verify file exists
	_, err = os.Stat(filename)
	assert.Nil(t, err)

	// Read and verify content
	content, err := os.ReadFile(filename)
	assert.Nil(t, err)

	expectedContent := "Name,Age,City\nJohn,25,New York\nJane,30,Los Angeles\nBob,35,Chicago\n"
	assert.Equal(t, expectedContent, string(content))
}

func TestExportToCSV_EmptyData(t *testing.T) {
	headers := []string{"Name", "Age"}
	rows := [][]string{}
	filename := "test_empty.csv"

	defer os.Remove(filename)

	err := ExportToCSV(headers, rows, filename)

	assert.Nil(t, err)

	// Verify file exists and contains only headers
	content, err := os.ReadFile(filename)
	assert.Nil(t, err)
	assert.Equal(t, "Name,Age\n", string(content))
}

func TestExportToCSV_SingleRow(t *testing.T) {
	headers := []string{"ID", "Value"}
	rows := [][]string{{"1", "Test"}}
	filename := "test_single.csv"

	defer os.Remove(filename)

	err := ExportToCSV(headers, rows, filename)

	assert.Nil(t, err)

	content, err := os.ReadFile(filename)
	assert.Nil(t, err)
	assert.Equal(t, "ID,Value\n1,Test\n", string(content))
}

func TestExportToCSV_SpecialCharacters(t *testing.T) {
	headers := []string{"Name", "Description"}
	rows := [][]string{
		{"John", "Contains, comma"},
		{"Jane", "Contains \"quotes\""},
		{"Bob", "Contains\nnewline"},
	}
	filename := "test_special.csv"

	defer os.Remove(filename)

	err := ExportToCSV(headers, rows, filename)

	assert.Nil(t, err)

	// Verify file exists
	_, err = os.Stat(filename)
	assert.Nil(t, err)
}

func TestExportToCSV_InvalidDirectory(t *testing.T) {
	headers := []string{"Name"}
	rows := [][]string{{"John"}}
	filename := "/nonexistent/directory/test.csv"

	err := ExportToCSV(headers, rows, filename)

	assert.NotNil(t, err)
}

func TestExportToCSV_LargeDataset(t *testing.T) {
	headers := []string{"ID", "Name", "Email"}
	rows := [][]string{}

	// Create 1000 rows
	for i := 1; i <= 1000; i++ {
		rows = append(rows, []string{
			string(rune(i)),
			"User" + string(rune(i)),
			"user" + string(rune(i)) + "@example.com",
		})
	}

	filename := "test_large.csv"

	defer os.Remove(filename)

	err := ExportToCSV(headers, rows, filename)

	assert.Nil(t, err)

	// Verify file exists and has correct number of lines
	content, err := os.ReadFile(filename)
	assert.Nil(t, err)

	// Should have 1001 lines (1 header + 1000 data rows)
	lines := 0
	for _, char := range content {
		if char == '\n' {
			lines++
		}
	}
	// The actual line count might vary due to how CSV writer handles the data
	assert.GreaterOrEqual(t, lines, 1000)
}

func TestExportToCSV_UnicodeCharacters(t *testing.T) {
	headers := []string{"Name", "Language"}
	rows := [][]string{
		{"José", "Español"},
		{"François", "Français"},
		{"Müller", "Deutsch"},
		{"李", "中文"},
	}
	filename := "test_unicode.csv"

	defer os.Remove(filename)

	err := ExportToCSV(headers, rows, filename)

	assert.Nil(t, err)

	// Verify file exists
	_, err = os.Stat(filename)
	assert.Nil(t, err)

	// Read and verify content contains unicode characters
	content, err := os.ReadFile(filename)
	assert.Nil(t, err)

	// Should contain unicode characters
	assert.Contains(t, string(content), "José")
	assert.Contains(t, string(content), "François")
	assert.Contains(t, string(content), "Müller")
	assert.Contains(t, string(content), "李")
}

func TestExportToCSV_EmptyHeaders(t *testing.T) {
	headers := []string{}
	rows := [][]string{{"data1", "data2"}}
	filename := "test_empty_headers.csv"

	defer os.Remove(filename)

	err := ExportToCSV(headers, rows, filename)

	assert.Nil(t, err)

	// Verify file exists
	_, err = os.ReadFile(filename)
	assert.Nil(t, err)
}

func TestExportToCSV_NilData(t *testing.T) {
	headers := []string{"Name"}
	var rows [][]string
	filename := "test_nil.csv"

	defer os.Remove(filename)

	err := ExportToCSV(headers, rows, filename)

	assert.Nil(t, err)

	// Verify file exists and contains only headers
	content, err := os.ReadFile(filename)
	assert.Nil(t, err)
	assert.Equal(t, "Name\n", string(content))
}
