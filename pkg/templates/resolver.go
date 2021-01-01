package templates

const resolverTpl = `
opt := req.ProtoReflect().Descriptor().Options()
ext := proto.GetExtension(opt, pb.E_Messages).(*pb.Messages)
resolver := ext.GetResolver()

if resolver == nil {
  return nil, status.Error(codes.InvalidArgument, "no resolver found")
}

config := map[string]interface{}{}

r, err := s.builder.Build(resolver.Name, config, resolver.Name)
if err != nil {
  return nil, status.Error(codes.Internal, err.Error())
}

rr, err := r()
if err != nil {
  return nil, status.Error(codes.Internal, err.Error())
}

var resp {{ name .Method }}_{{ .Method.Output.Name }}
if err := rr.Resolve(ctx, req, &resp); err != nil {
  return nil, status.Error(codes.Internal, err.Error())
}

return &resp, nil
`
