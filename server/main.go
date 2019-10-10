package main

import (
	gorpc_pb "gorpc/proto"
	"log"
)

func main() {
	request := &gorpc_pb.EchoRequest{
		Msg: "Varun",
	}

	log.Printf("%v", request)
}
