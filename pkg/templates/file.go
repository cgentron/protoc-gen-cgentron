package templates

const fileTpl = `// Code generated by aws-grpc-service-proxy. DO NOT EDIT.

package {{ pkg . }}

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"math"
	"net"
  "time"

  "github.com/andersnormal/pkg/server"
	o "github.com/cgentron/protoc-gen-cgentron/pkg/opts"
  "github.com/cgentron/protoc-gen-cgentron/pkg/proxy"
  "github.com/cgentron/protoc-gen-cgentron/pkg/resolvers"
  pb "github.com/cgentron/api/proto"

  "github.com/golang/protobuf/jsonpb"
  "google.golang.org/protobuf/proto"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	grpc_health_v1 "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
)

type srv struct {
	opts      *o.Opts
	resolvers o.ResolverRules
}

type service struct {
	tlsCfg  *tls.Config
	logger  *zap.Logger
	builder *resolvers.Builder
	UnimplementedExampleServer
}

func New(opts *o.Opts) proxy.Listener {
	s := new(srv)
	s.opts = opts

	s.resolvers = resolver_rules
	for k, r := range s.opts.Resolvers {
		s.resolvers[k] = r
	}

	return s
}

func NewProxy(opts *o.Opts) proxy.Proxy {
	s := New(opts)
	p := proxy.New(s, opts)

	return p
}

func (s *srv) Start(ctx context.Context, ready server.ReadyFunc) func() error {
	return func() error {
		lis, err := net.Listen("tcp", s.opts.Addr)
		if err != nil {
			return err
		}

		c, err := resolvers.NewClient()
		if err != nil {
			return err
		}

		rr, err := s.fetchResolvers(ctx, c)
		if err != nil {
			return err
		}

		b, err := resolvers.NewBuilder(c, rr)
		if err != nil {
			return err
		}

		ll := s.opts.Logger.With(zap.String("addr", s.opts.Addr))
		srv := &service{
			logger:  s.opts.Logger,
			builder: b,
		}

		tlsConfig := &tls.Config{}
		tlsConfig.InsecureSkipVerify = true
		srv.tlsCfg = tlsConfig

		var kaep = keepalive.EnforcementPolicy{
			MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
			PermitWithoutStream: true,            // Allow pings even when there are no active streams
		}

		var kasp = keepalive.ServerParameters{
			MaxConnectionIdle:     time.Duration(math.MaxInt64), // If a client is idle for 15 seconds, send a GOAWAY
			MaxConnectionAge:      time.Duration(math.MaxInt64), // If any connection is alive for more than 30 seconds, send a GOAWAY
			MaxConnectionAgeGrace: 5 * time.Second,              // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
			Time:                  5 * time.Second,              // Ping the client if it is idle for 5 seconds to ensure the connection is still active
			Timeout:               1 * time.Second,              // Wait 1 second for the ping ack before assuming the connection is dead
		}

		grpc_zap.ReplaceGrpcLogger(ll)

		ss := grpc.NewServer(
			grpc.KeepaliveEnforcementPolicy(kaep),
			grpc.KeepaliveParams(kasp),
			grpc_middleware.WithUnaryServerChain(grpc_zap.UnaryServerInterceptor(ll)))

		RegisterExampleServer(ss, srv)
		grpc_health_v1.RegisterHealthServer(ss, health.NewServer())
		// grpc_health_v1.RegisterHealthServer(ss, srv)

		ready()

		ll.Info("start listening")

		if err := ss.Serve(lis); err != nil {
			return err
		}

		return nil
	}
}

func (s *srv) fetchResolvers(ctx context.Context, c *resolvers.Client) (map[string]resolvers.Descriptor, error) {
	rr := map[string]resolvers.Descriptor{}

	for _, p := range s.resolvers { // todo: could be errgroup
		_, err := c.Fetch(ctx, p.Name, p.Version, p.Url)
		if err != nil {
			return rr, err
		}

		if err := c.Unzip(p.Name, p.Version); err != nil {
			return rr, err
		}

		rr[p.Name] = resolvers.Descriptor{ModuleName: p.Name, Version: p.Version}
	}

	return rr, nil
}`
