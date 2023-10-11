package mtype

import (
	"context"
	"io"
	"log"

	pb "github.com/nielvid/go-userservice-grpc/proto"
)

// bidirectional streaming i.e receive streams and send stream
func Chat(client pb.UserServiceClient) {

	stream, err := client.Chat(context.Background())

	if err != nil {
		log.Fatalf("request not sent: %v", err)
	}
	wg := make(chan struct{})
	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("message not received: %v", err)
			}
			log.Fatalf("message received: %v", res.Message)
			req := &pb.ChatMessage{Message: "How are you" }
			if err := stream.Send(req); err != nil {
			log.Fatalf("request not sent: %v", err)
		}

		}
		<- wg
		stream.CloseSend()

		close(wg)
	}()

}
