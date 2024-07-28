package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	templateDir  = "event_invitation_templates"
	templateFile = "event_mail_format.txt"
	csvFilePath  = "band_list/band_names.csv"
	outputDir    = "send_mail_template"
)

func main() {
	// Read the email template
	templatePath := fmt.Sprintf("%s/%s", templateDir, templateFile)
	template, err := ioutil.ReadFile(templatePath)
	if err != nil {
		fmt.Println("Error reading template file:", err)
		return
	}

	// Open the CSV file
	csvFile, err := os.Open(csvFilePath)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer csvFile.Close()

	// Read the CSV file
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Read() // Skip header row

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	// Create the output directory if it doesn't exist
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating output directory:", err)
		return
	}

	// Generate a text file for each band
	for _, record := range records {
		bandName := record[0]
		output := strings.ReplaceAll(string(template), "%%BAND_NAME%%", bandName)
		fileName := fmt.Sprintf("%s/%s.txt", outputDir, bandName)

		err = ioutil.WriteFile(fileName, []byte(output), 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
			return
		}

		fmt.Printf("Generated %s\n", fileName)
	}
}
