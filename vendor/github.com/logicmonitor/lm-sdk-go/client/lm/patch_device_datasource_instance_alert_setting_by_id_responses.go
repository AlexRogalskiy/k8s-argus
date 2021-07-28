// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/logicmonitor/lm-sdk-go/models"
)

// PatchDeviceDatasourceInstanceAlertSettingByIDReader is a Reader for the PatchDeviceDatasourceInstanceAlertSettingByID structure.
type PatchDeviceDatasourceInstanceAlertSettingByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchDeviceDatasourceInstanceAlertSettingByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchDeviceDatasourceInstanceAlertSettingByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPatchDeviceDatasourceInstanceAlertSettingByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchDeviceDatasourceInstanceAlertSettingByIDOK creates a PatchDeviceDatasourceInstanceAlertSettingByIDOK with default headers values
func NewPatchDeviceDatasourceInstanceAlertSettingByIDOK() *PatchDeviceDatasourceInstanceAlertSettingByIDOK {
	return &PatchDeviceDatasourceInstanceAlertSettingByIDOK{}
}

/*PatchDeviceDatasourceInstanceAlertSettingByIDOK handles this case with default header values.

successful operation
*/
type PatchDeviceDatasourceInstanceAlertSettingByIDOK struct {
	Payload *models.DeviceDataSourceInstanceAlertSetting
}

func (o *PatchDeviceDatasourceInstanceAlertSettingByIDOK) Error() string {
	return fmt.Sprintf("[PATCH /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{instanceId}/alertsettings/{id}][%d] patchDeviceDatasourceInstanceAlertSettingByIdOK  %+v", 200, o.Payload)
}

func (o *PatchDeviceDatasourceInstanceAlertSettingByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.DeviceDataSourceInstanceAlertSetting)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchDeviceDatasourceInstanceAlertSettingByIDDefault creates a PatchDeviceDatasourceInstanceAlertSettingByIDDefault with default headers values
func NewPatchDeviceDatasourceInstanceAlertSettingByIDDefault(code int) *PatchDeviceDatasourceInstanceAlertSettingByIDDefault {
	return &PatchDeviceDatasourceInstanceAlertSettingByIDDefault{
		_statusCode: code,
	}
}

/*PatchDeviceDatasourceInstanceAlertSettingByIDDefault handles this case with default header values.

Error
*/
type PatchDeviceDatasourceInstanceAlertSettingByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch device datasource instance alert setting by Id default response
func (o *PatchDeviceDatasourceInstanceAlertSettingByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchDeviceDatasourceInstanceAlertSettingByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{instanceId}/alertsettings/{id}][%d] patchDeviceDatasourceInstanceAlertSettingById default  %+v", o._statusCode, o.Payload)
}

func (o *PatchDeviceDatasourceInstanceAlertSettingByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
