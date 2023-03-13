// Write a Go program to parse a CSV file: The intern should write a Go program that reads a
// CSV file and parses its contents into a data structure (e.g., a slice of structs). The program
// should then output the data in a formatted way (e.g., as a table)

// Name - Jhalak Dashora
// Mail - jhalakdashora01@gmail.com

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
)

// defining a struct to hold the CSV data
type Record struct {
	Name   string
	Age    int
	Height float64
}

func main() {
	// opening the CSV file
	file, err := os.Open("data.csv")

	// checking if there is any error
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// creating a new CSV reader
	reader := csv.NewReader(file)

	var records []Record

	for {
		record, err := reader.Read()
		// if end of file the end the loop
		if err == io.EOF {
			break
		}
		// if any error then log it
		if err != nil {
			log.Fatal(err)
		}
		// fetching the data
		// getting age
		age, _ := strconv.Atoi(record[1])

		// getting height
		height, _ := strconv.ParseFloat(record[2], 64)

		// appending in record
		records = append(records, Record{
			Name:   record[0],
			Age:    age,
			Height: height,
		})
	}

	// printing the output in a formatted way
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "Name\tAge\tHeight")
	for _, r := range records {
		fmt.Fprintf(w, "%s\t%d\t%.2f\n", r.Name, r.Age, r.Height)
	}
	w.Flush()
}
