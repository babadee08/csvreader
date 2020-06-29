package csvwriter

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func Write(csvData [][]string, filename string) {
	recordFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file", err)
		return
	}

	writer := csv.NewWriter(recordFile)

	err = writer.WriteAll(csvData)

	if err != nil {
		fmt.Println("Error while writing to the file ::", err)
		return
	}

	err = recordFile.Close()
	if err != nil {
		fmt.Println("Error while closing the file ::", err)
		return
	}
}

func WriteLine(csvData [][]string, filename string) {
	recordFile, err := os.Create(filename)

	if err != nil {
		fmt.Println("Error Creating file", err)
		return
	}

	writer := csv.NewWriter(recordFile)

	for i, csvLineData := range csvData {
		if i == 0 {
			csvLineData = append(csvLineData, "Name Length")
		} else {
			nameLength := strconv.Itoa(len(csvLineData[0]))
			csvLineData = append(csvLineData, nameLength)
		}
		err = writer.Write(append(csvLineData, string(i)))
		if err != nil {
			fmt.Println("Error while writing to the file ::", err)
			return
		}

		writer.Flush()       // Writes the buffered data to the writer
		err = writer.Error() // Checks if any error occurred while writing
		if err != nil {
			fmt.Println("Error while writing to the file ::", err)
			return
		}
	}

	err = recordFile.Close()
	if err != nil {
		fmt.Println("Error while closing the file ::", err)
		return
	}
}
