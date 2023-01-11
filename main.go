package main

import (
	"flag"
	"log"
	"os"
	"simpleGRPC/checkconfig"
	"simpleGRPC/configstruct"
	"simpleGRPC/grpchub"
	"simpleGRPC/slurper"
)

//Protobuf
//protoc --go_out=. --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false protoraw/*.proto

func main() {
	log.Println("Simple gRPC v1.0")
	//cli flags for config file path
	configPath := flag.String("config", "./config.yml", "Path to config file")
	flag.Parse()

	//read config at path or create a default template
	slurper.Init(*configPath)

	//checkconfig
	checkconfig.Init()

	switch configstruct.CurrentConfig.SimpleGRPC.Settings.RunType {
	case "client":
		grpchub.StartClient()
	case "server":
		grpchub.StartServer()
	default:
		log.Println("No recognized run type")
		os.Exit(1)
	}
}
