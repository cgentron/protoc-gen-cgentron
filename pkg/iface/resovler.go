package iface

import (
	"context"
)

// Resolver ...
type Resolver interface {
	// Resolve ...
	Resolv(context.Context) (interface{}, error)
}
