// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	"golang.org/x/net/context"
)

// NewImportBatchJobParams creates a new ImportBatchJobParams object
// with the default values initialized.
func NewImportBatchJobParams() *ImportBatchJobParams {
	var ()
	return &ImportBatchJobParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewImportBatchJobParamsWithTimeout creates a new ImportBatchJobParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewImportBatchJobParamsWithTimeout(timeout time.Duration) *ImportBatchJobParams {
	var ()
	return &ImportBatchJobParams{

		timeout: timeout,
	}
}

// NewImportBatchJobParamsWithContext creates a new ImportBatchJobParams object
// with the default values initialized, and the ability to set a context for a request
func NewImportBatchJobParamsWithContext(ctx context.Context) *ImportBatchJobParams {
	var ()
	return &ImportBatchJobParams{

		Context: ctx,
	}
}

// NewImportBatchJobParamsWithHTTPClient creates a new ImportBatchJobParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewImportBatchJobParamsWithHTTPClient(client *http.Client) *ImportBatchJobParams {
	var ()
	return &ImportBatchJobParams{
		HTTPClient: client,
	}
}

/*ImportBatchJobParams contains all the parameters to send to the API endpoint
for the import batch job operation typically these are written to a http.Request
*/
type ImportBatchJobParams struct {

	/*File*/
	File runtime.NamedReadCloser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the import batch job params
func (o *ImportBatchJobParams) WithTimeout(timeout time.Duration) *ImportBatchJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the import batch job params
func (o *ImportBatchJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the import batch job params
func (o *ImportBatchJobParams) WithContext(ctx context.Context) *ImportBatchJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the import batch job params
func (o *ImportBatchJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the import batch job params
func (o *ImportBatchJobParams) WithHTTPClient(client *http.Client) *ImportBatchJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the import batch job params
func (o *ImportBatchJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFile adds the file to the import batch job params
func (o *ImportBatchJobParams) WithFile(file runtime.NamedReadCloser) *ImportBatchJobParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the import batch job params
func (o *ImportBatchJobParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WriteToRequest writes these params to a swagger request
func (o *ImportBatchJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form file param file
	if err := r.SetFileParam("file", o.File); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
