package main

import (
	"fmt"

	_ "github.com/golang/protobuf/jsonpb"
	_ "github.com/golang/protobuf/ptypes"
	_ "github.com/golang/protobuf/ptypes/any"
	_ "github.com/golang/protobuf/ptypes/wrappers"

	_ "github.com/wyattanderson/googleapis-reproducer/proto"
	_ "google.golang.org/protobuf/types/known/anypb"
)

func main() {
	fmt.Println("hello from legacy")
}
