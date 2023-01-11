package grpchub

import (
	"log"
	"net"
	"simpleGRPC/configstruct"
	"simpleGRPC/protocompiled"
	"strconv"

	"google.golang.org/grpc"
)

/*
	INSECURE SERVER START
*/

func StartServer() {
	uintPort := configstruct.CurrentConfig.SimpleGRPC.Settings.Listenport //port as unsigned integer
	lPort := strconv.Itoa(int(uintPort))                                  //cast uint to int; int to string ... I know it is retarded

	//open tcp socket listener for grpc
	lstn, lstnErr := net.Listen("tcp", "0.0.0.0:"+lPort)
	if lstnErr != nil {
		log.Fatal(lstnErr)
	}
	defer lstn.Close() //failsafe

	//new grpc server
	grpcServer := grpc.NewServer()

	//register grpc
	protocompiled.RegisterHelloServiceServer(grpcServer, &Server{}) //debug hello world grpc call

	//output for human readability
	log.Println("Listening on 0.0.0.0:" + lPort)

	//start grpc server on the tcp listener
	if grpcErr := grpcServer.Serve(lstn); grpcErr != nil {
		log.Fatal(grpcErr)
	}
}
