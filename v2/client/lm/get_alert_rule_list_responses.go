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

// GetAlertRuleListReader is a Reader for the GetAlertRuleList structure.
type GetAlertRuleListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertRuleListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAlertRuleListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetAlertRuleListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAlertRuleListOK creates a GetAlertRuleListOK with default headers values
func NewGetAlertRuleListOK() *GetAlertRuleListOK {
	return &GetAlertRuleListOK{}
}

/*GetAlertRuleListOK handles this case with default header values.

successful operation
*/
type GetAlertRuleListOK struct {
	Payload *models.AlertRulePaginationResponse
}

func (o *GetAlertRuleListOK) Error() string {
	return fmt.Sprintf("[GET /setting/alert/rules][%d] getAlertRuleListOK  %+v", 200, o.Payload)
}

func (o *GetAlertRuleListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertRulePaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertRuleListDefault creates a GetAlertRuleListDefault with default headers values
func NewGetAlertRuleListDefault(code int) *GetAlertRuleListDefault {
	return &GetAlertRuleListDefault{
		_statusCode: code,
	}
}

/*GetAlertRuleListDefault handles this case with default header values.

Error
*/
type GetAlertRuleListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get alert rule list default response
func (o *GetAlertRuleListDefault) Code() int {
	return o._statusCode
}

func (o *GetAlertRuleListDefault) Error() string {
	return fmt.Sprintf("[GET /setting/alert/rules][%d] getAlertRuleList default  %+v", o._statusCode, o.Payload)
}

func (o *GetAlertRuleListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}