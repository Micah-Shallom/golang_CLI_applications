package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

func sum(data []float64) float64 {
	sum := 0.0
	for _, v := range data {
		sum += v
	}
	return sum
}

func avg(data []float64) float64 {
	return sum(data) / float64(len(data))
}

func max(data []float64) float64{
	maxVal := data[0]
	for _, val := range data{
		if val > maxVal{
			maxVal = val
		}
	}
	return maxVal
}
func min(data []float64) float64{
	minVal := data[0]
	for _, val := range data{
		if val < minVal{
			minVal = val
		}
	}
	return minVal
}

// statsFunc define a generic statistical function
type statsFunc func(data []float64) float64

func csv2float(r io.Reader, column int) ([]float64, error) {
	var data []float64

	//create the CSV reader used to read in data  from csv files
	cr := csv.NewReader(r)
	cr.ReuseRecord = true //reuse same slice for each operation to reduce memory allocation

	//Adjusting for 0 based index
	column--

	// The method ReadAll() reads in all records (lines) from the CSV file as a slice of fields (columns), where each field is itself a slice of strings. Go represents this data structure as [][]string.

	// allData, err := cr.ReadAll()
	// if err != nil {
	// 	return nil, fmt.Errorf("cannot read data from file: %w", err)
	// }

	for i:=0; ; i++{
		row, err := cr.Read() //memory optimization procedure

		if err == io.EOF{
			break
		}
		if err != nil {
			return nil, fmt.Errorf("cannot read data from file: %w", err)
		}
		if i == 0 {
			continue
		}
		//Checking number of columns in CSV file
		if len(row) <= column {
			//file does not have many columns
			return nil, fmt.Errorf("%w: File has only %d columns", ErrInvalidColumn, len(row))
		}
		//try to convert data read into a float number
		v, err := strconv.ParseFloat(row[column], 64)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrNotNumber, err)
		}
		data = append(data, v)
	}
	return data, nil
}

	//looping through all records
	// for i, row := range allData {
	// 	if i == 0 {
	// 		continue
	// 	}
	// 	//Checking number of columns in CSV file
	// 	if len(row) <= column {
	// 		//file does not have many columns
	// 		return nil, fmt.Errorf("%w: File has only %d columns", ErrInvalidColumn, len(row))
	// 	}
	// 	//try to convert data read into a float number
	// 	v, err := strconv.ParseFloat(row[column], 64)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("%w: %s", ErrNotNumber, err)
	// 	}
	// 	data = append(data, v)
	// }
	// return data, nil

