package module

import (
	"bytes"
	"io"

	pb "github.com/cgentron/api/proto"
	"github.com/cgentron/protoc-gen-cgentron/pkg/templates"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

type Module struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
}

func Proxy() pgs.Module { return &Module{ModuleBase: &pgs.ModuleBase{}} }

func (m *Module) InitContext(ctx pgs.BuildContext) {
	m.ModuleBase.InitContext(ctx)
	m.ctx = pgsgo.InitContext(ctx.Parameters())
}

func (m *Module) Name() string { return "proxy" }

func (m *Module) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	buf := &bytes.Buffer{}

	for _, f := range targets {
		m.genProxy(f, buf, m.Parameters())
	}

	return m.Artifacts()
}

func (m *Module) genProxy(f pgs.File, buf *bytes.Buffer, params pgs.Parameters) {
	m.Push(f.Name().String())
	defer m.Pop()

	buf.Reset()
	v := initProxyVisitor(m, buf, "", params)
	m.CheckErr(pgs.Walk(v, f), "unable to generate AST tree")
	m.CheckErr(postProcessProxyVisitor(v, params), "unable to generate plugins")

	out := buf.String()

	m.AddGeneratorFile(
		m.ctx.OutputPath(f).SetExt(".proxy.go").String(),
		out,
	)
}

type ProxyVisitor struct {
	pgs.Visitor
	pgs.DebuggerCommon
	pgs.Parameters
	Resolvers templates.RegisteredResolvers
	prefix    string
	w         io.Writer
}

func initProxyVisitor(d pgs.DebuggerCommon, w io.Writer, prefix string, params pgs.Parameters) ProxyVisitor {
	p := ProxyVisitor{
		prefix:         prefix,
		w:              w,
		Parameters:     params,
		Resolvers:      make(templates.RegisteredResolvers),
		DebuggerCommon: d,
	}

	p.Visitor = pgs.PassThroughVisitor(&p)

	return p
}

func postProcessProxyVisitor(v ProxyVisitor, params pgs.Parameters) error {
	tpl, err := templates.Resolvers(v.Parameters, v.Resolvers)
	if err != nil {
		return err
	}

	err = tpl.Execute(v.w, params)
	if err != nil {
		return err
	}

	return nil
}

// VisitFile ...
func (p ProxyVisitor) VisitFile(f pgs.File) (pgs.Visitor, error) {
	tpl, err := templates.File(p.Parameters)
	if err != nil {
		return nil, err
	}

	err = tpl.Execute(p.w, f)
	if err != nil {
		return nil, err
	}

	return p, err
}

// VisitService ...
func (p ProxyVisitor) VisitService(s pgs.Service) (pgs.Visitor, error) {
	tpl, err := templates.Service(p.Parameters)
	if err != nil {
		return nil, err
	}

	err = tpl.Execute(p.w, s)
	if err != nil {
		return nil, err
	}

	return p, err
}

// VisitMethod ...
func (p ProxyVisitor) VisitMethod(m pgs.Method) (pgs.Visitor, error) {
	// if m.ServerStreaming() {
	// 	return p.visitMethodServerSideStreaming(m)
	// }

	// if m.ClientStreaming() {
	// 	return p.visitMethodClientSideStreaming(m)
	// }

	return p.visitMethod(m)
}

func (p ProxyVisitor) visitMethod(m pgs.Method) (pgs.Visitor, error) {
	tpl, err := templates.Method(p.Parameters)
	if err != nil {
		return nil, err
	}

	err = tpl.Execute(p.w, m)
	if err != nil {
		return nil, err
	}

	return p, err
}

// func (p ProxyVisitor) visitMethodServerSideStreaming(m pgs.Method) (pgs.Visitor, error) {
// 	tpl, err := templates.MethodServerStreaming(p.Parameters)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = tpl.Execute(p.w, m)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return p, err
// }

// func (p ProxyVisitor) visitMethodClientSideStreaming(m pgs.Method) (pgs.Visitor, error) {
// 	tpl, err := templates.MethodClientStreaming(p.Parameters)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = tpl.Execute(p.w, m)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return p, err
// }

// VisitMessage ...
func (p ProxyVisitor) VisitMessage(m pgs.Message) (pgs.Visitor, error) {
	if err := p.registerResolver(m); err != nil {
		return nil, err
	}

	tpl, err := templates.Message(p.Parameters)
	if err != nil {
		return nil, err
	}

	err = tpl.Execute(p.w, m)
	if err != nil {
		return nil, err
	}

	return p, err
}

func (p ProxyVisitor) registerResolver(m interface{}) error {
	var resolver *pb.ResolverRule

	switch ext := m.(type) {
	case pgs.Message:
		var mm pb.Messages
		if _, err := ext.Extension(pb.E_Messages, &mm); err != nil {
			return err
		}

		resolver = mm.GetResolver()
	default:
	}

	if resolver != nil {
		p.Debugf("found resolver: %s", resolver.Name)
		p.Resolvers[resolver.Name] = resolver
	}

	return nil
}

var _ pgs.Module = (*Module)(nil)