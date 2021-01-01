package templates

const resolversTpl = `var resolver_rules_rawDesc = map[string]*pb.ResolverRule{
  {{ range $name, $resolver := resolvers }}
    "{{ $name }}": {
      Name: "{{ $resolver.Name }}",
      Url: "{{ $resolver.Url }}",
      Version: "{{ $resolver.Version }}",
    },
  {{ end }}
}

var (
	resolver_rules o.ResolverRules = resolver_rules_rawDesc
)
`
