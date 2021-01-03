package templr

import (
	"text/template"

	"github.com/cgentron/protoc-gen-cgentron/pkg/iface"

	pgs "github.com/lyft/protoc-gen-star"
)

var _ iface.Termplr = (*templr)(nil)

// Opts ...
type Opts struct {
}

// Opt ...
type Opt func(*Opts)

// Configure ...
func (s *Opts) Configure(opts ...Opt) error {
	for _, o := range opts {
		o(s)
	}

	return nil
}

type templr struct {
	opts *Opts
}

// New ...
func New(opts ...Opt) iface.Termplr {
	options := new(Opts)
	options.Configure(opts...)

	t := new(templr)
	t.opts = options

	return t
}

func (t *templr) Make(templateType iface.TemplateType, tpl string, fn iface.RegisterFn, params pgs.Parameters) (*template.Template, error) {
	tt := template.New((string(templateType)))
	fn(tt, params)

	return tt.Parse(tpl)
}
