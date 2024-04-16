package httpx

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"os/exec"
)

func RunHttpx(domains []string) {
	writeTempFile(domains)
	cmd := exec.Command("httpx", "-l", "tempHttpx.txt", "-silent", "-td", "-csv", "-o", "temp.csv")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("  => httpx run completed")
	processCSV()
	os.Remove("tempHttpx.txt")
	os.Remove("temp.csv")
}

func writeTempFile(list []string) {
	file, err := os.OpenFile("tempHttpx.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
	}

	datawriter := bufio.NewWriter(file)

	for _, data := range list {
		_, _ = datawriter.WriteString(data + "\n")
	}

	datawriter.Flush()
	file.Close()
}

func processCSV() {
	// Open the CSV file
	inputFile, err := os.Open("temp.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	// Create a new CSV reader
	reader := csv.NewReader(inputFile)

	// Read the whole CSV into memory
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	// Specify the indices of the columns to keep
	columnsToKeep := []int{0, 1, 7, 8, 10, 11, 12, 17, 20, 26, 27, 28, 32, 33, 35, 37} // Keeping only the first and third columns (0-indexed)

	// Open the output CSV file
	outputFile, err := os.Create("EasyEASM.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Iterate through the records, keeping only the specified columns, and write to the output file
	for _, record := range records {
		var filteredRecord []string
		for _, columnIndex := range columnsToKeep {
			if columnIndex < len(record) {
				filteredRecord = append(filteredRecord, record[columnIndex])
			}
		}
		if err := writer.Write(filteredRecord); err != nil {
			fmt.Println(err)
			return
		}
	}

	// Ensure everything is written to the output file
	if err := writer.Error(); err != nil {
		fmt.Println(err)
		return
	}

}
