package main

import (
	"fmt"

	"github.com/ciscomsk/otus_golang_basic_2023_07_hw/hw02_fix_app/printer"
	"github.com/ciscomsk/otus_golang_basic_2023_07_hw/hw02_fix_app/reader"
	"github.com/ciscomsk/otus_golang_basic_2023_07_hw/hw02_fix_app/types"
)

func main() {
	var path string

	fmt.Println("Enter data file path: ")
	_, err := fmt.Scanln(&path)
	if err != nil {
		path = "data.json"
	}

	var staff []types.Employee

	staff, err = reader.ReadJSON(path)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	printer.PrintStaff(staff)
}
