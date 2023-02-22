package main

import (
	"fmt"

	_ "cloud.google.com/go/vision/apiv1"
	_ "google.golang.org/genproto/googleapis/cloud/vision/v1"
	_ "google.golang.org/grpc"

	_ "github.com/wyattanderson/googleapis-reproducer/proto"
)

func main() {
	fmt.Println("hello from v1")
}
