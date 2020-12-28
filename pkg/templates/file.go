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
  "io"

	"github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/lambda"
  "github.com/aws/aws-sdk-go/service/sqs"
	"github.com/golang/protobuf/jsonpb"
  "google.golang.org/grpc"
  "google.golang.org/grpc/health"
  "google.golang.org/grpc/keepalive"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
  grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
  grpc_health_v1 "google.golang.org/grpc/health/grpc_health_v1"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/service/dynamodb"
  "go.uber.org/zap"

  "github.com/cgentron/protoc-gen-cgentron/pkg/proxy"
  api "github.com/cgentron/protoc-gen-cgentron/api"
  o "github.com/cgentron/protoc-gen-cgentron/pkg/opts"
)
`
