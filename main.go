package main

import (
	"encoding/csv"
	"encoding/xml"
	"log"
	"os"
)

type String struct {
	Key   string `xml:"Key"`
	Value string `xml:"Value"`
}

type Entry struct {
	String []String `xml:"String"`
}

type Group struct {
	Name  string  `xml:"Name"`
	Entry []Entry `xml:"Entry"`
}

type Groups struct {
	Group []Group `xml:"Group"`
}

type Root struct {
	Groups []Groups `xml:"Group"`
}

type KeePass struct {
	Root Root `xml:"Root"`
}

func (r *KeePass) Log() {
	// Parse over all groups inside root.
	for _, gs := range r.Root.Groups {
		// Parse over all the groups inside the main group.
		for _, gsi := range gs.Group {
			// Parse over all entries.
			for _, en := range gsi.Entry {
				// Parse over all the strings inside the entry.
				for _, ens := range en.String {
					// Log all the fields inside.
					log.Println(ens.Key, ens.Value)
				}
			}
		}
	}
}

func (r *KeePass) Format() [][]string {
	// Find the length the csvSlice should be.
	var entryCount int
	for _, gs := range r.Root.Groups {
		for _, gsi := range gs.Group {
			for range gsi.Entry {
				// Increment csvSlice length counter by 1.
				entryCount++
			}
		}
	}

	// Add 1 more to entry count to account for the csvSlice format fields.
	entryCount++

	// Make a slice to store the XML data.
	csvSlice := make([][]string, entryCount)

	// KeePassXCs requierd fields (for CSV format):
	// Group,Title,Username,Password,URL,Notes,TOTP,Icon,Last Modified,Created
	csvSlice[0] = []string{
		"Group",
		"Title",
		"Username",
		"Password",
		"URL",
		"Notes",
		"TOTP",
		"Icon",
		"Last Modified",
		"Created",
	}

	// Create variable to keep track of how many entries have been added.
	var addedCount int

	// Update count since I added the csvSlice format fields (line 73).
	addedCount++

	// Parse over all groups inside root.
	for _, gs := range r.Root.Groups {
		// Parse over all the groups inside the main group.
		for _, gsi := range gs.Group {
			// Parse over all entries.
			for _, en := range gsi.Entry {
				entrySlice := make([]string, 10)
				// Add group name to entry slice.
				entrySlice[0] = gsi.Name

				// Parse over all the strings inside the entry.
				for _, ens := range en.String {
					// Add value depending on keys value.
					switch ens.Key {
					case "Title":
						entrySlice[1] = ens.Value
					case "UserName":
						entrySlice[2] = ens.Value
					case "Password":
						entrySlice[3] = ens.Value
					case "URL":
						entrySlice[4] = ens.Value
					case "Notes":
						entrySlice[5] = ens.Value
					}
				}
				// Add entry slice to csvSlice (the slice that will be returned as result).
				csvSlice[addedCount] = entrySlice

				// Update added entry count.
				addedCount++
			}
		}
	}

	// Return csvSlice as CSV formatted result.
	return csvSlice
}

func main() {
	// TODO: Add command-line input and multi-structure support
	fileName := "example"
	fileDir := "tmp/"
	// Open KeePassXC XML file.
	xmlData, err := os.ReadFile(fileDir + fileName + ".xml") // Read file.
	if err != nil {
		log.Fatal(err)
	}

	// Create a struct with the format for the XML data.
	var newData KeePass

	// Unmarshal the XML data into the struct.
	if err = xml.Unmarshal(xmlData, &newData); err != nil {
		log.Fatal(err)
	}

	// Log the XML data in a readable format.
	//newData.Log()

	// Create a CSV file.
	csvFile, err := os.Create(fileDir + fileName + ".csv")
	if err != nil {
		log.Fatal(err)
	}

	// Format the XML data into a CSV format.
	csvData := newData.Format()

	// Write formatted data to the csv file.
	writer := csv.NewWriter(csvFile)
	writer.WriteAll(csvData)

	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}
}
