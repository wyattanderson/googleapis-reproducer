package main

import (
	"fmt"

	_ "cloud.google.com/go/vision/v2/apiv1"
	_ "cloud.google.com/go/vision/v2/apiv1/visionpb"
	_ "google.golang.org/grpc"

	_ "github.com/wyattanderson/googleapis-reproducer/proto"
)

func main() {
	fmt.Println("hello from v2")
}
