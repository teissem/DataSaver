package main

import (
	"io"
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
	logFile, err := os.OpenFile(config.Log, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("[ERROR] Opening file : " + err.Error())
	}
	defer func() {
		err = logFile.Close()
		if err != nil {
			log.Printf("[ERROR] failed to close %s : %s", config.Log, err.Error())
		}
	}()
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
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
