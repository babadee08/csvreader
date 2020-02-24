package main

import (
	"fmt"
	"github.com/babadee08/csvreader/csvreader"
	"github.com/babadee08/csvreader/csvwriter"
)

var csvData = [][]string{
	{"SuperHero Name", "Power", "Weakness"},
	{"Batman", "Wealth", "Human"},
	{"Superman", "Strength", "Kryptonite"},
}

func main() {
	fmt.Println("CSV Reader sample")

	csvwriter.Write(csvData, "./superhero.csv")
	csvwriter.WriteLine(csvData, "./superheros2.csv")
	csvreader.ReadAll("records.csv")
	csvreader.ReadRow("records.csv")
}
