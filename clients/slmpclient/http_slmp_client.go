package slmpclient

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/SAP/cf-mta-plugin/clients/slmpclient/operations"
)

// Default slmp HTTP client.
var Default = NewHTTPClient(nil)

// NewHTTPClient creates a new slmp HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Slmp {
	if formats == nil {
		formats = strfmt.Default
	}
	transport := httptransport.New("example.com", "/slprot/slmp/", []string{"http", "https"})
	return New(transport, formats)
}

// New creates a new slmp client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Slmp {
	cli := new(Slmp)
	cli.Transport = transport

	cli.Operations = operations.New(transport, formats)

	return cli
}

// Slmp is a client for slmp
type Slmp struct {
	Operations *operations.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Slmp) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Operations.SetTransport(transport)

}