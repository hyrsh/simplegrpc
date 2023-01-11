package configstruct

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

//CurrentConfig stores our config at runtime (can also be updated during runtime)
var CurrentConfig Config

//CurrentConfigPath ... pretty self-explanatory
var CurrentConfigPath string

//Config sets the pattern for our config
type Config struct {
	SimpleGRPC struct {
		Settings struct {
			Listenport uint16 `yaml:"listenport"`
			Interval   string `yaml:"interval"`
			Target     string `yaml:"target"`
			Message    string `yaml:"message"`
			Answer     string `yaml:"answer"`
			RunType    string `yaml:"runtype"`
		} `yaml:"settings"`
	} `yaml:"simple-grpc"`
}

//SetConfig sets the config for central access (and partial updating (port changes still require a restart))
func SetConfig(config Config) {
	CurrentConfig = config
}

//SetConfigPath ... and again (gets called from configchanger.go)
func SetConfigPath(configpath string) {
	CurrentConfigPath = configpath
}

//ConfigWriter only gets called when no config is found to provide a template
func ConfigWriter(file string) {
	template := Config{}
	//set default values
	writeDefaults(&template)
	//marshal interface to byte array
	output, outputErr := yaml.Marshal(&template)
	if outputErr != nil {
		log.Println("YAML marshal error!")
		log.Println(outputErr)
	}
	//make sure the path exists and then write template to file
	filePath := filepath.Dir(file)
	pErr := os.MkdirAll(filePath, 0755)
	if pErr != nil {
		log.Println("Cannot create path" + string(filePath))
		log.Println(pErr)
	}
	//write data to file
	writeErr := ioutil.WriteFile(file, output, 0755)
	if writeErr != nil {
		log.Println("YAML cannot write data to file!")
		log.Println(writeErr)
	}
}

//kind of self-explanatory. We set default values since this stupid struct is not able to do this on its own
func writeDefaults(config *Config) {
	//Settings
	config.SimpleGRPC.Settings.Listenport = 8080
	config.SimpleGRPC.Settings.Interval = "1s"
	config.SimpleGRPC.Settings.Target = "127.0.0.1:9090"
	config.SimpleGRPC.Settings.Message = "Hi from a simple gRPC client!"
	config.SimpleGRPC.Settings.Answer = "Hi from a simple gRPC server!"
	config.SimpleGRPC.Settings.RunType = "client"
}
