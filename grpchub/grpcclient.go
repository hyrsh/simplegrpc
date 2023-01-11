package grpchub

import (
	"context"
	"log"
	"simpleGRPC/configstruct"
	"simpleGRPC/protocompiled"
	"time"

	"google.golang.org/grpc"
)

//the client first checks some basic functionality
func StartClient() {
	dur, _ := time.ParseDuration(configstruct.CurrentConfig.SimpleGRPC.Settings.Intervall) //error check happens earlier
	for {
		sendHello(configstruct.CurrentConfig.SimpleGRPC.Settings.Target)
		time.Sleep(dur)
	}
}

func sendHello(ep string) {
	call, callErr := grpc.Dial(ep, grpc.WithInsecure()) //create connection object
	if callErr != nil {
		log.Println("gRPC call error!")
		log.Println(callErr.Error())
	}
	defer call.Close() //failsafe

	//register new client service on connection object
	rq := protocompiled.NewHelloServiceClient(call)

	//Sssend the snake for echo
	msg := protocompiled.Message{
		Message: configstruct.CurrentConfig.SimpleGRPC.Settings.Message,
	}

	//rpc call context with timeout per call
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond) //RETRY POSSIBILITY
	defer cancel()

	//get rpc server response
	ret, retErr := rq.Hello(ctx, &msg)
	if retErr != nil {
		log.Println("Cannot reach target " + configstruct.CurrentConfig.SimpleGRPC.Settings.Target)
	}
	if ret != nil {
		log.Println(ret.Answer)
	}
}
