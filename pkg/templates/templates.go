package templates

import (
	"text/template"

	pb "github.com/cgentron/api/proto"
	"github.com/cgentron/protoc-gen-cgentron/pkg/templates/shared"

	pgs "github.com/lyft/protoc-gen-star"
)

type RegisterFn func(tpl *template.Template, params pgs.Parameters)
type RegisteredResolvers map[string]*pb.ResolverRule

func makeTemplate(ext string, fn RegisterFn, params pgs.Parameters) *template.Template {
	tpl := template.New(ext)
	shared.Register(tpl, params)

	template.Must(tpl.New("resolver").Parse(resolverTpl))

	fn(tpl, params)
	return tpl
}

func File(params pgs.Parameters) (*template.Template, error) {
	tpl := makeTemplate("file", func(tpl *template.Template, params pgs.Parameters) {}, params)
	return tpl.Parse(fileTpl)
}

func Service(params pgs.Parameters) (*template.Template, error) {
	tpl := makeTemplate("service", func(tpl *template.Template, params pgs.Parameters) {}, params)
	return tpl.Parse(serviceTpl)
}

func Method(params pgs.Parameters) (*template.Template, error) {
	tpl := makeTemplate("method", func(tpl *template.Template, params pgs.Parameters) {}, params)
	return tpl.Parse(methodTpl)
}

func Message(params pgs.Parameters) (*template.Template, error) {
	tpl := makeTemplate("message", func(tpl *template.Template, params pgs.Parameters) {}, params)
	return tpl.Parse(messageTpl)
}

func Resolvers(params pgs.Parameters, resolvers RegisteredResolvers) (*template.Template, error) {
	tpl := makeTemplate("resolvers", func(tpl *template.Template, params pgs.Parameters) {
		tpl.Funcs(map[string]interface {
		}{
			"resolvers": func() RegisteredResolvers { return resolvers },
		})
	}, params)
	return tpl.Parse(resolversTpl)
}
