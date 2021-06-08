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

// GetDataSourceOverviewGraphListReader is a Reader for the GetDataSourceOverviewGraphList structure.
type GetDataSourceOverviewGraphListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDataSourceOverviewGraphListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDataSourceOverviewGraphListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetDataSourceOverviewGraphListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDataSourceOverviewGraphListOK creates a GetDataSourceOverviewGraphListOK with default headers values
func NewGetDataSourceOverviewGraphListOK() *GetDataSourceOverviewGraphListOK {
	return &GetDataSourceOverviewGraphListOK{}
}

/*GetDataSourceOverviewGraphListOK handles this case with default header values.

successful operation
*/
type GetDataSourceOverviewGraphListOK struct {
	Payload *models.DatasourceOverviewGraphPaginationResponse
}

func (o *GetDataSourceOverviewGraphListOK) Error() string {
	return fmt.Sprintf("[GET /setting/datasources/{dsId}/ographs][%d] getDataSourceOverviewGraphListOK  %+v", 200, o.Payload)
}

func (o *GetDataSourceOverviewGraphListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.DatasourceOverviewGraphPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataSourceOverviewGraphListDefault creates a GetDataSourceOverviewGraphListDefault with default headers values
func NewGetDataSourceOverviewGraphListDefault(code int) *GetDataSourceOverviewGraphListDefault {
	return &GetDataSourceOverviewGraphListDefault{
		_statusCode: code,
	}
}

/*GetDataSourceOverviewGraphListDefault handles this case with default header values.

Error
*/
type GetDataSourceOverviewGraphListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get data source overview graph list default response
func (o *GetDataSourceOverviewGraphListDefault) Code() int {
	return o._statusCode
}

func (o *GetDataSourceOverviewGraphListDefault) Error() string {
	return fmt.Sprintf("[GET /setting/datasources/{dsId}/ographs][%d] getDataSourceOverviewGraphList default  %+v", o._statusCode, o.Payload)
}

func (o *GetDataSourceOverviewGraphListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
