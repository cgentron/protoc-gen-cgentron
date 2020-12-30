// Code generated by aws-grpc-service-proxy. DO NOT EDIT.

package proto

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"math"
	"net"
	"time"

	"github.com/andersnormal/pkg/server"
	pb "github.com/cgentron/api/proto"
	o "github.com/cgentron/protoc-gen-cgentron/pkg/opts"
	"github.com/cgentron/protoc-gen-cgentron/pkg/proxy"
	"github.com/cgentron/protoc-gen-cgentron/pkg/resolvers"

	"github.com/golang/protobuf/jsonpb"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	grpc_health_v1 "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
)

var resolver_rules_rawDesc = map[string]*pb.ResolverRule{
	"": {},
}

var (
	resolver_rules o.ResolverRules = resolver_rules_rawDesc
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
}

// SongJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of SongJSONMarshaler This struct is safe to replace or modify but
// should not be done so concurrently.
var SongJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Song) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := SongJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Song)(nil)

// SongJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Song. This struct is safe to replace or modify but
// should not be done so concurrently.
var SongJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Song) UnmarshalJSON(b []byte) error {
	return SongJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Song)(nil)

// InsertJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of InsertJSONMarshaler This struct is safe to replace or modify but
// should not be done so concurrently.
var InsertJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Insert) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := InsertJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Insert)(nil)

// InsertJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Insert. This struct is safe to replace or modify but
// should not be done so concurrently.
var InsertJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Insert) UnmarshalJSON(b []byte) error {
	return InsertJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Insert)(nil)

// Insert_RequestJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Insert_RequestJSONMarshaler This struct is safe to replace or modify but
// should not be done so concurrently.
var Insert_RequestJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Insert_Request) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := Insert_RequestJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Insert_Request)(nil)

// Insert_RequestJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Insert_Request. This struct is safe to replace or modify but
// should not be done so concurrently.
var Insert_RequestJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Insert_Request) UnmarshalJSON(b []byte) error {
	return Insert_RequestJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Insert_Request)(nil)

// Insert_ResponseJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Insert_ResponseJSONMarshaler This struct is safe to replace or modify but
// should not be done so concurrently.
var Insert_ResponseJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Insert_Response) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := Insert_ResponseJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Insert_Response)(nil)

// Insert_ResponseJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Insert_Response. This struct is safe to replace or modify but
// should not be done so concurrently.
var Insert_ResponseJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Insert_Response) UnmarshalJSON(b []byte) error {
	return Insert_ResponseJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Insert_Response)(nil)

// Here goes a message Insert
func (s *service) Insert(ctx context.Context, req *Insert_Request) (*Insert_Response, error) {

	// resolver
	return nil, nil

}