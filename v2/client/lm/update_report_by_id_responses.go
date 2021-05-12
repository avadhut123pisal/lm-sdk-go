// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/logicmonitor/lm-sdk-go/v2/models"
)

// UpdateReportByIDReader is a Reader for the UpdateReportByID structure.
type UpdateReportByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateReportByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateReportByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateReportByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateReportByIDOK creates a UpdateReportByIDOK with default headers values
func NewUpdateReportByIDOK() *UpdateReportByIDOK {
	return &UpdateReportByIDOK{}
}

/*UpdateReportByIDOK handles this case with default header values.

successful operation
*/
type UpdateReportByIDOK struct {
	Payload models.ReportBase
}

func (o *UpdateReportByIDOK) Error() string {
	return fmt.Sprintf("[PUT /report/reports/{id}][%d] updateReportByIdOK  %+v", 200, o.Payload)
}

func (o *UpdateReportByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalReportBase(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewUpdateReportByIDDefault creates a UpdateReportByIDDefault with default headers values
func NewUpdateReportByIDDefault(code int) *UpdateReportByIDDefault {
	return &UpdateReportByIDDefault{
		_statusCode: code,
	}
}

/*UpdateReportByIDDefault handles this case with default header values.

Error
*/
type UpdateReportByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update report by Id default response
func (o *UpdateReportByIDDefault) Code() int {
	return o._statusCode
}

func (o *UpdateReportByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /report/reports/{id}][%d] updateReportById default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateReportByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
