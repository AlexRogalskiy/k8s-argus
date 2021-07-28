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

// GetWebsiteGroupByIDReader is a Reader for the GetWebsiteGroupByID structure.
type GetWebsiteGroupByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWebsiteGroupByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetWebsiteGroupByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetWebsiteGroupByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWebsiteGroupByIDOK creates a GetWebsiteGroupByIDOK with default headers values
func NewGetWebsiteGroupByIDOK() *GetWebsiteGroupByIDOK {
	return &GetWebsiteGroupByIDOK{}
}

/*GetWebsiteGroupByIDOK handles this case with default header values.

successful operation
*/
type GetWebsiteGroupByIDOK struct {
	Payload *models.WebsiteGroup
}

func (o *GetWebsiteGroupByIDOK) Error() string {
	return fmt.Sprintf("[GET /website/groups/{id}][%d] getWebsiteGroupByIdOK  %+v", 200, o.Payload)
}

func (o *GetWebsiteGroupByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.WebsiteGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebsiteGroupByIDDefault creates a GetWebsiteGroupByIDDefault with default headers values
func NewGetWebsiteGroupByIDDefault(code int) *GetWebsiteGroupByIDDefault {
	return &GetWebsiteGroupByIDDefault{
		_statusCode: code,
	}
}

/*GetWebsiteGroupByIDDefault handles this case with default header values.

Error
*/
type GetWebsiteGroupByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get website group by Id default response
func (o *GetWebsiteGroupByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetWebsiteGroupByIDDefault) Error() string {
	return fmt.Sprintf("[GET /website/groups/{id}][%d] getWebsiteGroupById default  %+v", o._statusCode, o.Payload)
}

func (o *GetWebsiteGroupByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
