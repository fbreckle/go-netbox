go-netbox
=========

[![Go Reference](https://pkg.go.dev/badge/github.com/fbreckle/go-netbox.svg)](https://pkg.go.dev/github.com/fbreckle/go-netbox) [![Report Card](https://goreportcard.com/badge/github.com/fbreckle/go-netbox)](https://goreportcard.com/report/github.com/fbreckle/go-netbox)

The `netbox/client` and `netbox/models` packages are a Go client for the REST API of
[netbox-community's NetBox](https://github.com/netbox-community/netbox) IPAM and DCIM service,
generated from NetBox's own OpenAPI document, kept in this repository.

Why this fork exists
====================

This fork exists solely to support [e-breuninger/terraform-provider-netbox](https://github.com/e-breuninger/terraform-provider-netbox). As such, some changes in this fork do only make sense in that context.

Goals
=====

NetBox has published OpenAPI 3 since 3.5 ([announcement](https://github.com/netbox-community/netbox/discussions/11808)),
and go-swagger — which generates this client — reads OpenAPI 2 only. The spec is converted from
version 3 down to version 2 rather than switching generators, as netbox-community's go-netbox did
([issue](https://github.com/netbox-community/go-netbox/issues/155),
[discussion](https://github.com/netbox-community/go-netbox/discussions/156)): I prefer the code
go-swagger produces to what the OpenAPI 3 generators emit, and that preference is why I maintain
this fork.

How the spec is maintained
==========================

`tools/v3migrate/migrate.py` converts the v3 spec to v2 spec
`tools/v3migrate/README.md` documents the tool and its translation rules, and
`tools/v3migrate/patches.py` holds what the conversion cannot derive, each entry with the reason it
is there.

Versioning
==========

tbd. Meanwhile, look at branches and tags.

Using the client
================

The client is a go-swagger client: build a transport for your NetBox host, put the API token on it,
and every operation can then be called with `nil` for `authInfo`. For example:

```go
package main

import (
	"log"
	"os"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	httptransport "github.com/go-openapi/runtime/client"
)

func main() {
	token := os.Getenv("NETBOX_TOKEN")
	if token == "" {
		log.Fatalf("Please provide netbox API token via env var NETBOX_TOKEN")
	}

	netboxHost := os.Getenv("NETBOX_HOST")
	if netboxHost == "" {
		log.Fatalf("Please provide netbox host via env var NETBOX_HOST")
	}

	transport := httptransport.New(netboxHost, client.DefaultBasePath, []string{"https"})
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "Token "+token)

	c := client.New(transport, nil)

	req := dcim.NewDcimSitesListParams()
	res, err := c.Dcim.DcimSitesList(req, nil)
	if err != nil {
		log.Fatalf("Cannot get sites list: %v", err)
	}
	log.Printf("res: %v", res)
}
```

Go Module support
================

Go 1.25+

`go get github.com/fbreckle/go-netbox`


More complex client configuration
=================================

The client is generated using [go-swagger](https://github.com/go-swagger/go-swagger). This means the generated client
makes use of [github.com/go-openapi/runtime/client](https://godoc.org/github.com/go-openapi/runtime/client). If you need
a more complex configuration, it is probably possible with a combination of this generated client and the runtime
options.

The [godocs for the go-openapi/runtime/client module](https://godoc.org/github.com/go-openapi/runtime/client) explain
the client options in detail, including different authentication and debugging options. Worth knowing: setting the
`DEBUG` environment variable dumps every request to standard out.

Regenerating the client
=======================

`make generate` converts the tracked v3 document to `swagger.processed.json` and runs go-swagger on
it; `make check` does that from a clean tree and fails if the tracked client changed — that is the
whole drift test.

Moving to a newer NetBox version means replacing the document with that version's, downloaded from
`/api/schema/`, then regenerating and committing both:

```
make clean generate
```
