package grpchub

import (
	"context"
	"log"
	"simpleGRPC/configstruct"
	"simpleGRPC/protocompiled"
)

//empty server struct for grpc <-- this is necessary for later processing
type Server struct {
}

/*
	GRPC FUNCTIONS
*/

//default hello world grpc call
func (s *Server) Hello(ctx context.Context, message *protocompiled.Message) (*protocompiled.Answer, error) { //error maybe used but I am not sure yet
	log.Println(message.Message)
	return &protocompiled.Answer{Answer: configstruct.CurrentConfig.SimpleGRPC.Settings.Answer}, nil
}
