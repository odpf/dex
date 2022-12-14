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

// NewGetProjectBySlugParams creates a new GetProjectBySlugParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProjectBySlugParams() *GetProjectBySlugParams {
	return &GetProjectBySlugParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectBySlugParamsWithTimeout creates a new GetProjectBySlugParams object
// with the ability to set a timeout on a request.
func NewGetProjectBySlugParamsWithTimeout(timeout time.Duration) *GetProjectBySlugParams {
	return &GetProjectBySlugParams{
		timeout: timeout,
	}
}

// NewGetProjectBySlugParamsWithContext creates a new GetProjectBySlugParams object
// with the ability to set a context for a request.
func NewGetProjectBySlugParamsWithContext(ctx context.Context) *GetProjectBySlugParams {
	return &GetProjectBySlugParams{
		Context: ctx,
	}
}

// NewGetProjectBySlugParamsWithHTTPClient creates a new GetProjectBySlugParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProjectBySlugParamsWithHTTPClient(client *http.Client) *GetProjectBySlugParams {
	return &GetProjectBySlugParams{
		HTTPClient: client,
	}
}

/*
GetProjectBySlugParams contains all the parameters to send to the API endpoint

	for the get project by slug operation.

	Typically these are written to a http.Request.
*/
type GetProjectBySlugParams struct {

	/* Slug.

	   Unique slug of the project.
	*/
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get project by slug params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectBySlugParams) WithDefaults() *GetProjectBySlugParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get project by slug params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectBySlugParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get project by slug params
func (o *GetProjectBySlugParams) WithTimeout(timeout time.Duration) *GetProjectBySlugParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get project by slug params
func (o *GetProjectBySlugParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get project by slug params
func (o *GetProjectBySlugParams) WithContext(ctx context.Context) *GetProjectBySlugParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get project by slug params
func (o *GetProjectBySlugParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get project by slug params
func (o *GetProjectBySlugParams) WithHTTPClient(client *http.Client) *GetProjectBySlugParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get project by slug params
func (o *GetProjectBySlugParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSlug adds the slug to the get project by slug params
func (o *GetProjectBySlugParams) WithSlug(slug string) *GetProjectBySlugParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the get project by slug params
func (o *GetProjectBySlugParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectBySlugParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param slug
	if err := r.SetPathParam("slug", o.Slug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
