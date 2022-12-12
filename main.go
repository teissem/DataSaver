package main

import (
	"log"
	"os"
	"path"

	"golang.org/x/exp/maps"
	"teissem.fr/data_saver/src/compression"
	"teissem.fr/data_saver/src/configuration"
	"teissem.fr/data_saver/src/datasource"
)

func main() {
	// Arguments waited : <configuration_file>
	const waitedArguments = 2
	arguments := os.Args
	if len(arguments) < waitedArguments {
		log.Fatalf("[ERROR] Usage : ./main <configuration_file>")
	}
	// Verification of the compatibility of the configuration file
	fileExtension := path.Ext(os.Args[1])
	supportedConfigurationFormat := configuration.SupportedConfigurationFormat()
	confParser, ok := supportedConfigurationFormat[fileExtension]
	if !ok {
		log.Fatalf("[ERROR] Configuration file must be in format %v, current format is %s\n",
			maps.Keys(supportedConfigurationFormat),
			fileExtension)
	}
	// Parsing of the configuration
	config, err := confParser(os.Args[1])
	if err != nil {
		log.Fatalf("[ERROR] Parsing the configuration file : " + err.Error())
	}
	// Get all the data selected in the configuration file
	log.Printf("[INFO] Get data... ")
	err = datasource.GetData(config)
	if err != nil {
		log.Fatalf("[ERROR] Get data from source : " + err.Error())
	}
	log.Printf("[INFO] Get data done")
	// Compress the result
	log.Printf("[INFO] Compression... ")
	err = compression.Compress(config)
	if err != nil {
		log.Fatalf("[ERROR] Compression : " + err.Error())
	}
	log.Printf("[INFO] Compression done")
	log.Printf("[INFO] Data saved successfully")
}
