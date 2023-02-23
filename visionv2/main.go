package main

import (
	"fmt"
	"time"

	_ "cloud.google.com/go/vision/v2/apiv1"
	_ "cloud.google.com/go/vision/v2/apiv1/visionpb"
	_ "google.golang.org/grpc"

	"github.com/wyattanderson/googleapis-reproducer/proto"
	"google.golang.org/protobuf/types/known/durationpb"
)

func main() {
	fmt.Println("hello from v2")

	_ = &proto.GoogleCloudVision{
		Duration: durationpb.New(1 * time.Second),
	}
}
