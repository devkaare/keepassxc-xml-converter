package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
)

type Field struct {
	Key   string `xml:"Key"`
	Value string `xml:"Value"`
}

type Entry struct {
	Notes, Password, Title, URL, UserName string
}

func Decode(d *xml.Decoder) []Entry {
	var (
		result []Entry
		new    Entry
		field  Field
	)

	// Parse xml data
	for t, _ := d.Token(); t != nil; t, _ = d.Token() {
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "String" {
				d.DecodeElement(&field, &se)

				// Add x value to new
				switch field.Key {
				case "Notes":
					new.Notes = field.Value
				case "Password":
					new.Password = field.Value
				case "Title":
					new.Title = field.Value
				case "URL":
					new.URL = field.Value
				case "UserName":
					new.UserName = field.Value
					// UserName is the last field for an entry,
					// so new is appended to result when UserName is found
					result = append(result, new)
				}
			}
		}
	}

	// Return all result found
	return result
}

func Format(entries []Entry) [][]string {
	csvSlice := make([][]string, 1000) // Make a slice to store the XML data.

	var i int

	csvSlice[i] = []string{ // KeePassXCs requierd fields (for CSV format):
		"Group",
		"Title",
		"Username",
		"Password",
		"URL",
		"Notes",
	}
	i++

	for _, entry := range entries { // Add new entries
		csvSlice[i] = []string{
			"Root", // Root is the group entries will be stored in
			entry.Title,
			entry.UserName,
			entry.Password,
			entry.URL,
			entry.Notes,
		}
		i++
	}

	return csvSlice
}

func main() {
	fileDir := "tmp/"
	var fileName string
	fmt.Println("What's the name of the XML file you want to convert? e.g. passwords")
	fmt.Scanln(&fileName)

	fmt.Println("Opening file...")
	xmlFile, err := os.Open(filepath.Join(fileDir, (fileName + ".xml")))
	if err != nil {
		panic("Failed to open XML file")
	}

	// Decode
	fmt.Println("Decoding entries...")
	d := xml.NewDecoder(xmlFile)
	entries := Decode(d)

	// Create a CSV file
	fmt.Println("Writing output to file...")
	csvFile, err := os.Create(filepath.Join(fileDir, (fileName + ".csv")))
	if err != nil {
		panic("Failed to open CSV file")
	}

	// Format the XML data into a CSV format
	csvData := Format(entries)

	// Write formatted data to the csv file
	writer := csv.NewWriter(csvFile)
	writer.WriteAll(csvData)

	if err := writer.Error(); err != nil {
		panic("Failed to write to CSV file")
	}

	fmt.Println("Successfully converted XML file to CSV file!")
}
