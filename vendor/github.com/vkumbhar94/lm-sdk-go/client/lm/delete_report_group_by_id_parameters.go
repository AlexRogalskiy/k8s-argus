// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteReportGroupByIDParams creates a new DeleteReportGroupByIDParams object
// with the default values initialized.
func NewDeleteReportGroupByIDParams() *DeleteReportGroupByIDParams {
	var ()
	return &DeleteReportGroupByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteReportGroupByIDParamsWithTimeout creates a new DeleteReportGroupByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteReportGroupByIDParamsWithTimeout(timeout time.Duration) *DeleteReportGroupByIDParams {
	var ()
	return &DeleteReportGroupByIDParams{

		timeout: timeout,
	}
}

// NewDeleteReportGroupByIDParamsWithContext creates a new DeleteReportGroupByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteReportGroupByIDParamsWithContext(ctx context.Context) *DeleteReportGroupByIDParams {
	var ()
	return &DeleteReportGroupByIDParams{

		Context: ctx,
	}
}

// NewDeleteReportGroupByIDParamsWithHTTPClient creates a new DeleteReportGroupByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteReportGroupByIDParamsWithHTTPClient(client *http.Client) *DeleteReportGroupByIDParams {
	var ()
	return &DeleteReportGroupByIDParams{
		HTTPClient: client,
	}
}

/*DeleteReportGroupByIDParams contains all the parameters to send to the API endpoint
for the delete report group by Id operation typically these are written to a http.Request
*/
type DeleteReportGroupByIDParams struct {

	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete report group by Id params
func (o *DeleteReportGroupByIDParams) WithTimeout(timeout time.Duration) *DeleteReportGroupByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete report group by Id params
func (o *DeleteReportGroupByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete report group by Id params
func (o *DeleteReportGroupByIDParams) WithContext(ctx context.Context) *DeleteReportGroupByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete report group by Id params
func (o *DeleteReportGroupByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete report group by Id params
func (o *DeleteReportGroupByIDParams) WithHTTPClient(client *http.Client) *DeleteReportGroupByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete report group by Id params
func (o *DeleteReportGroupByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete report group by Id params
func (o *DeleteReportGroupByIDParams) WithID(id int32) *DeleteReportGroupByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete report group by Id params
func (o *DeleteReportGroupByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteReportGroupByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}