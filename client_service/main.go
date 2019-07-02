package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/snake/test/framework/serialization/amf"
	proto "github.com/snake/test/greeter"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("tcp.client"))
	service.Init()

	// Create new greeter client
	greeter := proto.NewTransceiverService("tcp", service.Client())

	// Call the greeter

	bodyBs, err := amf.EncodeAmf(123)
	rsp, err := greeter.Transmission(context.TODO(), &proto.SendRequest{MessageId: 1, Body: bodyBs})
	if err != nil {
		fmt.Println(err)
	}
	//binary
	// Print response
	fmt.Println(rsp.Body)
	fmt.Println(amf.Decode(bytes.NewBuffer(rsp.Body)))
}
