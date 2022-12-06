package main

import (
	"fmt"
	"os"

	"teissem.fr/data_saver/src/configuration"
)

func main() {
	// Arguments waited : <configuration_file>
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("Usage : ./main <configuration_file>")
		os.Exit(1)
	}
	config, err := configuration.ParseJSON(os.Args[1])
	if err != nil {
		fmt.Println("Error occurred when parsing JSON : " + err.Error())
	}
	fmt.Printf("%+v", config)
}
