package main

import (
	"fmt"

	api "github.com/cgentron/api/proto"
	pb "github.com/cgentron/pluginamzn/proto"
	example "github.com/cgentron/protoc-gen-cgentron/example"
	"google.golang.org/protobuf/proto"
)

func main() {
	m := &example.Insert_Request{}

	mopt := m.ProtoReflect().Descriptor().Options()
	mext := proto.GetExtension(mopt, pb.E_Messages).(*pb.Messages)
	m2ext := proto.GetExtension(mopt, api.E_Messages).(*api.Messages)

	fmt.Println(m2ext.GetResolver())
	fmt.Println(mext.GetLambda())
}
