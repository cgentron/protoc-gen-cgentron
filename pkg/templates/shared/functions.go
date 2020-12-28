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

	var services pb.Methods
	if _, err := m.Extension(pb.E_Methods, &services); err != nil {
		return ctx, err
	}

	ctx.Typ = resolveMethods(&services, m)
	ctx.Methods = &services

	if ctx.Typ == "error" {
		return ctx, fmt.Errorf("unknown template type")
	}

	return ctx, nil
}

func resolveMethods(a *pb.Methods, m pgs.Method) string {
	return "error"
}
