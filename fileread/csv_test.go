package fileread

import "testing"

func TestCsvFile(t *testing.T) {
	CsvReadByline("output.csv")
	CsvWriteByline("output.csv")
}
