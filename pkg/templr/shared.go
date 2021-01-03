package templr

import (
	"bytes"
	"fmt"
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

type SharedContext struct {
	Method pgs.Method

	Typ string
}

func Render(tpl *template.Template) func(ctx SharedContext) (string, error) {
	return func(ctx SharedContext) (string, error) {
		buf := &bytes.Buffer{}
		err := tpl.ExecuteTemplate(buf, ctx.Typ, ctx)
		return buf.String(), err
	}
}

type sharedFuncs struct {
	pgsgo.Context
}

func registerSharedFunc(tpl *template.Template, params pgs.Parameters) {
	fns := sharedFuncs{pgsgo.InitContext(params)}

	tpl.Funcs(map[string]interface {
	}{
		"pkg":     fns.PackageName,
		"name":    fns.Name,
		"context": sharedContext,
		"render":  Render(tpl),
	})
}

func sharedContext(m pgs.Method) (SharedContext, error) {
	ctx := SharedContext{}
	ctx.Method = m

	if ctx.Typ == "error" {
		return ctx, fmt.Errorf("unknown template type")
	}

	return ctx, nil
}
