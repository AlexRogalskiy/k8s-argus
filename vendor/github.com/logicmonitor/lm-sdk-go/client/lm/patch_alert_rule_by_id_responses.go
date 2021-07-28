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

// PatchAlertRuleByIDReader is a Reader for the PatchAlertRuleByID structure.
type PatchAlertRuleByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAlertRuleByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchAlertRuleByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPatchAlertRuleByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchAlertRuleByIDOK creates a PatchAlertRuleByIDOK with default headers values
func NewPatchAlertRuleByIDOK() *PatchAlertRuleByIDOK {
	return &PatchAlertRuleByIDOK{}
}

/*PatchAlertRuleByIDOK handles this case with default header values.

successful operation
*/
type PatchAlertRuleByIDOK struct {
	Payload *models.AlertRule
}

func (o *PatchAlertRuleByIDOK) Error() string {
	return fmt.Sprintf("[PATCH /setting/alert/rules/{id}][%d] patchAlertRuleByIdOK  %+v", 200, o.Payload)
}

func (o *PatchAlertRuleByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.AlertRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAlertRuleByIDDefault creates a PatchAlertRuleByIDDefault with default headers values
func NewPatchAlertRuleByIDDefault(code int) *PatchAlertRuleByIDDefault {
	return &PatchAlertRuleByIDDefault{
		_statusCode: code,
	}
}

/*PatchAlertRuleByIDDefault handles this case with default header values.

Error
*/
type PatchAlertRuleByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch alert rule by Id default response
func (o *PatchAlertRuleByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchAlertRuleByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /setting/alert/rules/{id}][%d] patchAlertRuleById default  %+v", o._statusCode, o.Payload)
}

func (o *PatchAlertRuleByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
