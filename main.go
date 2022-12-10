package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"golang.org/x/exp/maps"
	"teissem.fr/data_saver/src/configuration"
)

func main() {
	// Arguments waited : <configuration_file>
	const waitedArguments = 2
	arguments := os.Args
	if len(arguments) < waitedArguments {
		log.Fatalf("[ERROR] Usage : ./main <configuration_file>")
	}
	fileExtension := path.Ext(os.Args[1])
	supportedConfigurationFormat := configuration.SupportedConfigurationFormat()
	confParser, ok := supportedConfigurationFormat[fileExtension]
	if !ok {
		log.Fatalf("[ERROR] Configuration file must be in format %v, current format is %s\n",
			maps.Keys(supportedConfigurationFormat),
			fileExtension)
	}
	config, err := confParser(os.Args[1])
	if err != nil {
		log.Fatalf("[ERROR] JSON parsing : " + err.Error())
	}
	fmt.Printf("%+v", config)
}
