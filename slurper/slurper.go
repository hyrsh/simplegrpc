package slurper

import (
	"io/ioutil"
	"log"
	"os"
	"simpleGRPC/configstruct"
	"simpleGRPC/filehandling"

	"gopkg.in/yaml.v3"
)

//Init just consumes a path to a file
func Init(file string) {
	if filehandling.StatFile(file) {
		log.Println("Config found at " + file)
	} else {
		log.Println("Config not found. Creating template at " + file)
		configstruct.ConfigWriter(file)
		log.Println("You cannot start with a default config!")
		log.Println("Look into " + file + " and adjust it to your environment!")
		os.Exit(0)
	}
	loadConfig(file)
}

func loadConfig(file string) {
	configData, configError := ioutil.ReadFile(file)
	if configError != nil {
		log.Println("Could not read config file!")
		log.Println(configError)
	}
	var rawYAML configstruct.Config
	ymlErr := yaml.Unmarshal(configData, &rawYAML)
	if ymlErr != nil {
		log.Println("Unmarshal error!")
		log.Println(ymlErr)
	}
	configstruct.SetConfig(rawYAML)
	configstruct.SetConfigPath(file)
}
