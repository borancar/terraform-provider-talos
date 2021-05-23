module github.com/borancar/terraform-provider-talos

go 1.15

replace (
	github.com/talos-systems/talos => /home/boran/git/talos

	github.com/talos-systems/talos/pkg/machinery => /home/boran/git/talos/pkg/machinery

	// forked go-yaml that introduces RawYAML interface, which can be used to populate YAML fields using bytes
	// which are then encoded as a valid YAML blocks with proper indentiation
	gopkg.in/yaml.v3 => github.com/unix4ever/yaml v0.0.0-20210315173758-8fb30b8e5a5b

	// See https://github.com/talos-systems/go-loadbalancer/pull/4
	// `go get github.com/smira/tcpproxy@combined-fixes`, then copy pseudo-version there
	inet.af/tcpproxy => github.com/smira/tcpproxy v0.0.0-20201015133617-de5f7797b95b
)

require (
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.6.1
	github.com/talos-systems/talos v0.0.0-00010101000000-000000000000
	github.com/talos-systems/talos/pkg/machinery v0.0.0-20210302191918-8ffb55943c71
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)
