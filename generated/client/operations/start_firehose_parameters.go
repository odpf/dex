// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewStartFirehoseParams creates a new StartFirehoseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStartFirehoseParams() *StartFirehoseParams {
	return &StartFirehoseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStartFirehoseParamsWithTimeout creates a new StartFirehoseParams object
// with the ability to set a timeout on a request.
func NewStartFirehoseParamsWithTimeout(timeout time.Duration) *StartFirehoseParams {
	return &StartFirehoseParams{
		timeout: timeout,
	}
}

// NewStartFirehoseParamsWithContext creates a new StartFirehoseParams object
// with the ability to set a context for a request.
func NewStartFirehoseParamsWithContext(ctx context.Context) *StartFirehoseParams {
	return &StartFirehoseParams{
		Context: ctx,
	}
}

// NewStartFirehoseParamsWithHTTPClient creates a new StartFirehoseParams object
// with the ability to set a custom HTTPClient for a request.
func NewStartFirehoseParamsWithHTTPClient(client *http.Client) *StartFirehoseParams {
	return &StartFirehoseParams{
		HTTPClient: client,
	}
}

/*
StartFirehoseParams contains all the parameters to send to the API endpoint

	for the start firehose operation.

	Typically these are written to a http.Request.
*/
type StartFirehoseParams struct {

	// Body.
	Body interface{}

	/* FirehoseUrn.

	   URN of the firehose.
	*/
	FirehoseUrn string

	/* ProjectSlug.

	   Identifier for the project.
	*/
	ProjectSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the start firehose params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartFirehoseParams) WithDefaults() *StartFirehoseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the start firehose params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartFirehoseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the start firehose params
func (o *StartFirehoseParams) WithTimeout(timeout time.Duration) *StartFirehoseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start firehose params
func (o *StartFirehoseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start firehose params
func (o *StartFirehoseParams) WithContext(ctx context.Context) *StartFirehoseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start firehose params
func (o *StartFirehoseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start firehose params
func (o *StartFirehoseParams) WithHTTPClient(client *http.Client) *StartFirehoseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start firehose params
func (o *StartFirehoseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the start firehose params
func (o *StartFirehoseParams) WithBody(body interface{}) *StartFirehoseParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the start firehose params
func (o *StartFirehoseParams) SetBody(body interface{}) {
	o.Body = body
}

// WithFirehoseUrn adds the firehoseUrn to the start firehose params
func (o *StartFirehoseParams) WithFirehoseUrn(firehoseUrn string) *StartFirehoseParams {
	o.SetFirehoseUrn(firehoseUrn)
	return o
}

// SetFirehoseUrn adds the firehoseUrn to the start firehose params
func (o *StartFirehoseParams) SetFirehoseUrn(firehoseUrn string) {
	o.FirehoseUrn = firehoseUrn
}

// WithProjectSlug adds the projectSlug to the start firehose params
func (o *StartFirehoseParams) WithProjectSlug(projectSlug string) *StartFirehoseParams {
	o.SetProjectSlug(projectSlug)
	return o
}

// SetProjectSlug adds the projectSlug to the start firehose params
func (o *StartFirehoseParams) SetProjectSlug(projectSlug string) {
	o.ProjectSlug = projectSlug
}

// WriteToRequest writes these params to a swagger request
func (o *StartFirehoseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param firehoseUrn
	if err := r.SetPathParam("firehoseUrn", o.FirehoseUrn); err != nil {
		return err
	}

	// path param projectSlug
	if err := r.SetPathParam("projectSlug", o.ProjectSlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
