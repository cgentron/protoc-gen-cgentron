package shared

import (
	"fmt"
	"text/template"

	pb "github.com/cgentron/api/proto"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

type sharedFuncs struct {
	pgsgo.Context
}

func Register(tpl *template.Template, params pgs.Parameters) {
	fns := sharedFuncs{pgsgo.InitContext(params)}

	tpl.Funcs(map[string]interface {
	}{
		"pkg":     fns.PackageName,
		"name":    fns.Name,
		"context": proxyContext,
		"render":  Render(tpl),
	})
}

func proxyContext(m pgs.Method) (ProxyContext, error) {
	ctx := ProxyContext{}
	ctx.Method = m

	in := m.Input()
	var options pb.Messages
	if _, err := in.Extension(pb.E_Messages, &options); err != nil {
		return ctx, err
	}

	if options.GetResolver() != nil {
		ctx.Typ = "resolver"
	}

	// ctx.Typ = resolveMethods(&services, m)
	// ctx.Methods = &services

	if ctx.Typ == "error" {
		return ctx, fmt.Errorf("unknown template type")
	}

	return ctx, nil
}

func resolveMethods(a *pb.Methods, m pgs.Method) string {
	// println(a.GetResolver())
	// if a.GetResolver() != nil {
	// 	return "resolver"
	// }

	return "error"
}
