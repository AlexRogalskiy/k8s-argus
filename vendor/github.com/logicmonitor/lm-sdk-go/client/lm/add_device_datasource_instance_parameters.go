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
	"github.com/go-openapi/swag"
	models "github.com/logicmonitor/lm-sdk-go/models"
	"golang.org/x/net/context"
)

// NewAddDeviceDatasourceInstanceParams creates a new AddDeviceDatasourceInstanceParams object
// with the default values initialized.
func NewAddDeviceDatasourceInstanceParams() *AddDeviceDatasourceInstanceParams {
	var ()
	return &AddDeviceDatasourceInstanceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddDeviceDatasourceInstanceParamsWithTimeout creates a new AddDeviceDatasourceInstanceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddDeviceDatasourceInstanceParamsWithTimeout(timeout time.Duration) *AddDeviceDatasourceInstanceParams {
	var ()
	return &AddDeviceDatasourceInstanceParams{

		timeout: timeout,
	}
}

// NewAddDeviceDatasourceInstanceParamsWithContext creates a new AddDeviceDatasourceInstanceParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddDeviceDatasourceInstanceParamsWithContext(ctx context.Context) *AddDeviceDatasourceInstanceParams {
	var ()
	return &AddDeviceDatasourceInstanceParams{

		Context: ctx,
	}
}

// NewAddDeviceDatasourceInstanceParamsWithHTTPClient creates a new AddDeviceDatasourceInstanceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddDeviceDatasourceInstanceParamsWithHTTPClient(client *http.Client) *AddDeviceDatasourceInstanceParams {
	var ()
	return &AddDeviceDatasourceInstanceParams{
		HTTPClient: client,
	}
}

/*AddDeviceDatasourceInstanceParams contains all the parameters to send to the API endpoint
for the add device datasource instance operation typically these are written to a http.Request
*/
type AddDeviceDatasourceInstanceParams struct {

	/*Body*/
	Body *models.DeviceDataSourceInstance
	/*DeviceID*/
	DeviceID int32
	/*HdsID
	  The device-datasource ID

	*/
	HdsID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add device datasource instance params
func (o *AddDeviceDatasourceInstanceParams) WithTimeout(timeout time.Duration) *AddDeviceDatasourceInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add device datasource instance params
func (o *AddDeviceDatasourceInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add device datasource instance params
func (o *AddDeviceDatasourceInstanceParams) WithContext(ctx context.Context) *AddDeviceDatasourceInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add device datasource instance params
func (o *AddDeviceDatasourceInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add device datasource instance params
func (o *AddDeviceDatasourceInstanceParams) WithHTTPClient(client *http.Client) *AddDeviceDatasourceInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add device datasource instance params
func (o *AddDeviceDatasourceInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add device datasource instance params
func (o *AddDeviceDatasourceInstanceParams) WithBody(body *models.DeviceDataSourceInstance) *AddDeviceDatasourceInstanceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add device datasource instance params
func (o *AddDeviceDatasourceInstanceParams) SetBody(body *models.DeviceDataSourceInstance) {
	o.Body = body
}

// WithDeviceID adds the deviceID to the add device datasource instance params
func (o *AddDeviceDatasourceInstanceParams) WithDeviceID(deviceID int32) *AddDeviceDatasourceInstanceParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the add device datasource instance params
func (o *AddDeviceDatasourceInstanceParams) SetDeviceID(deviceID int32) {
	o.DeviceID = deviceID
}

// WithHdsID adds the hdsID to the add device datasource instance params
func (o *AddDeviceDatasourceInstanceParams) WithHdsID(hdsID int32) *AddDeviceDatasourceInstanceParams {
	o.SetHdsID(hdsID)
	return o
}

// SetHdsID adds the hdsId to the add device datasource instance params
func (o *AddDeviceDatasourceInstanceParams) SetHdsID(hdsID int32) {
	o.HdsID = hdsID
}

// WriteToRequest writes these params to a swagger request
func (o *AddDeviceDatasourceInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param deviceId
	if err := r.SetPathParam("deviceId", swag.FormatInt32(o.DeviceID)); err != nil {
		return err
	}

	// path param hdsId
	if err := r.SetPathParam("hdsId", swag.FormatInt32(o.HdsID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
