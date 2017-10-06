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

// NewGetProcessServiceParams creates a new GetProcessServiceParams object
// with the default values initialized.
func NewGetProcessServiceParams() *GetProcessServiceParams {
	var ()
	return &GetProcessServiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProcessServiceParamsWithTimeout creates a new GetProcessServiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProcessServiceParamsWithTimeout(timeout time.Duration) *GetProcessServiceParams {
	var ()
	return &GetProcessServiceParams{

		timeout: timeout,
	}
}

// NewGetProcessServiceParamsWithContext creates a new GetProcessServiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProcessServiceParamsWithContext(ctx context.Context) *GetProcessServiceParams {
	var ()
	return &GetProcessServiceParams{

		Context: ctx,
	}
}

/*GetProcessServiceParams contains all the parameters to send to the API endpoint
for the get process service operation typically these are written to a http.Request
*/
type GetProcessServiceParams struct {

	/*ProcessID*/
	ProcessID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the get process service params
func (o *GetProcessServiceParams) WithTimeout(timeout time.Duration) *GetProcessServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get process service params
func (o *GetProcessServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get process service params
func (o *GetProcessServiceParams) WithContext(ctx context.Context) *GetProcessServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get process service params
func (o *GetProcessServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithProcessID adds the processID to the get process service params
func (o *GetProcessServiceParams) WithProcessID(processID string) *GetProcessServiceParams {
	o.SetProcessID(processID)
	return o
}

// SetProcessID adds the processId to the get process service params
func (o *GetProcessServiceParams) SetProcessID(processID string) {
	o.ProcessID = processID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProcessServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param processId
	if err := r.SetPathParam("processId", o.ProcessID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}