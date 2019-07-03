package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	proto "github.com/snake/test/greeter"
)

type Transceiver struct{}

func (g *Transceiver) Transmission(ctx context.Context, req *proto.SendRequest, rsp *proto.SendResponse) error {
	rsp.MessageId = 1
	rsp.Body = req.Body
	rsp.Type = proto.MessageType_Notify

	fmt.Println(rsp.GetBody())

	return nil
}

func Auth() error {
	//

	return nil
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("tcp"),
	)

	// Init will parse the command line flags.
	service.Init()
	// Register handler
	_ = proto.RegisterTransceiverHandler(service.Server(), new(Transceiver))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
