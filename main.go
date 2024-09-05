package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Employee represents the structure of our JSON data
type Employee struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Department string `json:"department"`
}

// Employees is a collection of Employee
type Employees struct {
	Employees []Employee `json:"employees"`
}

func main() {
	// Open and read the JSON file
	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Parse the JSON data
	var employees Employees
	json.Unmarshal(byteValue, &employees)

	// Output the data
	fmt.Println("Employee Details:")
	for _, employee := range employees.Employees {
		fmt.Printf("Name: %s, Age: %d, Department: %s\n", employee.Name, employee.Age, employee.Department)
	}
}
