package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/xuri/excelize/v2"
)

type ServiceDate struct {
	ServiceDateID string `json:"service_date_id"`
	Action        string `json:"action"`
}

type Product struct {
	LicensedProductID string        `json:"licensed_product_id"`
	ServiceDates      []ServiceDate `json:"service_dates"`
}

type ProductData struct {
	ID       int `json:"id"`
	Products struct {
		Items []Product `json:"items"`
	} `json:"products"`
}

func main() {
	excelFile := "/Users/Rohit/Personal Development/Learning/GO Lang/MathWorks-Utilities/Payload Builders/licensed-product-service-date-removal-payload-data.xlsx" // Path to your Excel file
	sheetName := "ML-0000"                                                                                                                                         // Sheet name

	// Open the Excel file
	f, err := excelize.OpenFile(excelFile)
	if err != nil {
		log.Fatalf("Failed to open Excel file: %v", err)
	}
	defer f.Close()

	// Read all rows from the sheet
	rows, err := f.GetRows(sheetName)
	if err != nil {
		log.Fatalf("Failed to read rows from sheet: %v", err)
	}

	if len(rows) < 2 {
		log.Fatal("No data found in the Excel sheet.")
	}

	// Initialize ProductData structure
	data := ProductData{ID: 1234} // Set the ID as needed

	// Skip the header row and process the data
	for _, row := range rows[1:] {
		if len(row) < 2 {
			log.Println("Skipping row with insufficient columns")
			continue
		}

		product := Product{
			LicensedProductID: row[0],
			ServiceDates: []ServiceDate{
				{
					ServiceDateID: row[1],
					Action:        "REMOVE",
				},
			},
		}

		data.Products.Items = append(data.Products.Items, product)
	}

	// Convert to JSON
	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	// Write JSON to file
	outputFile := "output.json"
	if err := os.WriteFile(outputFile, jsonData, 0644); err != nil {
		log.Fatalf("Failed to write JSON to file: %v", err)
	}

	fmt.Printf("JSON has been written to %s\n", outputFile)
}
