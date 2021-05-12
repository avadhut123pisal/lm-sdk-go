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

// UpdateRecipientGroupByIDReader is a Reader for the UpdateRecipientGroupByID structure.
type UpdateRecipientGroupByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRecipientGroupByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateRecipientGroupByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateRecipientGroupByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateRecipientGroupByIDOK creates a UpdateRecipientGroupByIDOK with default headers values
func NewUpdateRecipientGroupByIDOK() *UpdateRecipientGroupByIDOK {
	return &UpdateRecipientGroupByIDOK{}
}

/*UpdateRecipientGroupByIDOK handles this case with default header values.

successful operation
*/
type UpdateRecipientGroupByIDOK struct {
	Payload *models.RecipientGroup
}

func (o *UpdateRecipientGroupByIDOK) Error() string {
	return fmt.Sprintf("[PUT /setting/recipientgroups/{id}][%d] updateRecipientGroupByIdOK  %+v", 200, o.Payload)
}

func (o *UpdateRecipientGroupByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecipientGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRecipientGroupByIDDefault creates a UpdateRecipientGroupByIDDefault with default headers values
func NewUpdateRecipientGroupByIDDefault(code int) *UpdateRecipientGroupByIDDefault {
	return &UpdateRecipientGroupByIDDefault{
		_statusCode: code,
	}
}

/*UpdateRecipientGroupByIDDefault handles this case with default header values.

Error
*/
type UpdateRecipientGroupByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update recipient group by Id default response
func (o *UpdateRecipientGroupByIDDefault) Code() int {
	return o._statusCode
}

func (o *UpdateRecipientGroupByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /setting/recipientgroups/{id}][%d] updateRecipientGroupById default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateRecipientGroupByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}