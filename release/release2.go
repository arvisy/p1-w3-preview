package release

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func Release2() {
	// Baca masukan CLI
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go input.csv output.csv")
		os.Exit(1)
	}

	inputFileName := os.Args[1]
	outputFileName := os.Args[2]

	// Baca File Input
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println("Error opening input file: ", err)
		os.Exit(1)
	}
	defer inputFile.Close()

	// Baca File Ouput
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Error opening output file: ", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	// Baca File CSV
	reader := csv.NewReader(inputFile)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		os.Exit(1)
	}

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Menulis header
	writer.Write(records[0])
	writer.Flush()

	// Pemrosesan setiap baris data
	for _, record := range records[1:] {
		processRecord(record, writer)
	}
}

func processRecord(data []string, writer *csv.Writer) {
	// Terapkan logika pemrosesan data di sini
	name := strings.ToUpper(data[0])
	age := strings.TrimSpace(data[1])
	occupation := "Mr. " + strings.TrimSpace(data[2])

	// Tulis data yang telah diproses ke file CSV output
	writer.Write([]string{name, age, occupation})
	writer.Flush()
}
