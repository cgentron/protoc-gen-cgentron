package iface

import (
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
)

// Template ...
type TemplateType string

// RegisterFn ..
type RegisterFn func(tpl *template.Template, params pgs.Parameters)

const (
	File    TemplateType = "file"
	Service TemplateType = "service"
	Method  TemplateType = "method"
	Message TemplateType = "message"
)

// Templr ...
type Termplr interface {
	Make(TemplateType, string, RegisterFn, pgs.Parameters) (*template.Template, error)
}
