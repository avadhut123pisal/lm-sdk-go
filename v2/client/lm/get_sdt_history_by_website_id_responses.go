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

// GetSDTHistoryByWebsiteIDReader is a Reader for the GetSDTHistoryByWebsiteID structure.
type GetSDTHistoryByWebsiteIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSDTHistoryByWebsiteIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSDTHistoryByWebsiteIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetSDTHistoryByWebsiteIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSDTHistoryByWebsiteIDOK creates a GetSDTHistoryByWebsiteIDOK with default headers values
func NewGetSDTHistoryByWebsiteIDOK() *GetSDTHistoryByWebsiteIDOK {
	return &GetSDTHistoryByWebsiteIDOK{}
}

/*GetSDTHistoryByWebsiteIDOK handles this case with default header values.

successful operation
*/
type GetSDTHistoryByWebsiteIDOK struct {
	Payload *models.WebsiteSDTHistoryPaginationResponse
}

func (o *GetSDTHistoryByWebsiteIDOK) Error() string {
	return fmt.Sprintf("[GET /website/websites/{id}/historysdts][%d] getSdtHistoryByWebsiteIdOK  %+v", 200, o.Payload)
}

func (o *GetSDTHistoryByWebsiteIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebsiteSDTHistoryPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSDTHistoryByWebsiteIDDefault creates a GetSDTHistoryByWebsiteIDDefault with default headers values
func NewGetSDTHistoryByWebsiteIDDefault(code int) *GetSDTHistoryByWebsiteIDDefault {
	return &GetSDTHistoryByWebsiteIDDefault{
		_statusCode: code,
	}
}

/*GetSDTHistoryByWebsiteIDDefault handles this case with default header values.

Error
*/
type GetSDTHistoryByWebsiteIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get SDT history by website Id default response
func (o *GetSDTHistoryByWebsiteIDDefault) Code() int {
	return o._statusCode
}

func (o *GetSDTHistoryByWebsiteIDDefault) Error() string {
	return fmt.Sprintf("[GET /website/websites/{id}/historysdts][%d] getSDTHistoryByWebsiteId default  %+v", o._statusCode, o.Payload)
}

func (o *GetSDTHistoryByWebsiteIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
