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

//On Windows
//$env:CGO_ENABLED = 1; $env:GOOS = "windows"; $env:GOARCH = "amd64"; go build -o simplegrpc_win_x86_64.exe -ldflags='-s -w -extldflags "-static"' main.go

//On Windows cross-compile for Linux
//$env:CGO_ENABLED = 0; $env:GOOS = "linux"; $env:GOARCH = "amd64"; go build -o simplegrpc_linux_x86_64 -ldflags='-s -w -extldflags "-static"' main.go

//On Linux (important for docker image building with "FROM scratch")
//go build -a -tags netgo --ldflags '-extldflags "-static"'

//packed with
//upx --best simplegrpc.exe (Windows)
//or
//upx --best simplegrpc (Linux)
//brute compression throws a false positive with Windows Defender -.-"

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
