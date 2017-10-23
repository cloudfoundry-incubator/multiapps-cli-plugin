package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetCurrentBreakpointsParams creates a new GetCurrentBreakpointsParams object
// with the default values initialized.
func NewGetCurrentBreakpointsParams() *GetCurrentBreakpointsParams {

	return &GetCurrentBreakpointsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCurrentBreakpointsParamsWithTimeout creates a new GetCurrentBreakpointsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCurrentBreakpointsParamsWithTimeout(timeout time.Duration) *GetCurrentBreakpointsParams {

	return &GetCurrentBreakpointsParams{

		timeout: timeout,
	}
}

// NewGetCurrentBreakpointsParamsWithContext creates a new GetCurrentBreakpointsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCurrentBreakpointsParamsWithContext(ctx context.Context) *GetCurrentBreakpointsParams {

	return &GetCurrentBreakpointsParams{

		Context: ctx,
	}
}

/*GetCurrentBreakpointsParams contains all the parameters to send to the API endpoint
for the get current breakpoints operation typically these are written to a http.Request
*/
type GetCurrentBreakpointsParams struct {
	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the get current breakpoints params
func (o *GetCurrentBreakpointsParams) WithTimeout(timeout time.Duration) *GetCurrentBreakpointsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get current breakpoints params
func (o *GetCurrentBreakpointsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get current breakpoints params
func (o *GetCurrentBreakpointsParams) WithContext(ctx context.Context) *GetCurrentBreakpointsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get current breakpoints params
func (o *GetCurrentBreakpointsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WriteToRequest writes these params to a swagger request
func (o *GetCurrentBreakpointsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}