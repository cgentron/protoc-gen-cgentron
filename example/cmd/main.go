package main

import (
	"context"
	"math/rand"
	"time"

	pb "github.com/cgentron/protoc-gen-cgentron/example"
	o "github.com/cgentron/protoc-gen-cgentron/pkg/opts"

	"go.uber.org/zap"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	opts := o.New(o.WithVerbose(), o.WithLogger(logger))
	p := pb.NewProxy(opts)

	if err := p.Start(ctx); err != nil {
		panic(err)
	}

	// m := &example.Insert_Request{}

	// mopt := m.ProtoReflect().Descriptor().Options()
	// mext := proto.GetExtension(mopt, pb.E_Messages).(*pb.Messages)
	// m2ext := proto.GetExtension(mopt, api.E_Messages).(*api.Messages)

	// fmt.Println(m2ext.GetResolver())
	// fmt.Println(mext.GetLambda())
}
