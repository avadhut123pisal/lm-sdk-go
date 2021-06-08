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

// PatchDevicePropertyByNameReader is a Reader for the PatchDevicePropertyByName structure.
type PatchDevicePropertyByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchDevicePropertyByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchDevicePropertyByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPatchDevicePropertyByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchDevicePropertyByNameOK creates a PatchDevicePropertyByNameOK with default headers values
func NewPatchDevicePropertyByNameOK() *PatchDevicePropertyByNameOK {
	return &PatchDevicePropertyByNameOK{}
}

/*PatchDevicePropertyByNameOK handles this case with default header values.

successful operation
*/
type PatchDevicePropertyByNameOK struct {
	Payload *models.EntityProperty
}

func (o *PatchDevicePropertyByNameOK) Error() string {
	return fmt.Sprintf("[PATCH /device/devices/{deviceId}/properties/{name}][%d] patchDevicePropertyByNameOK  %+v", 200, o.Payload)
}

func (o *PatchDevicePropertyByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.EntityProperty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchDevicePropertyByNameDefault creates a PatchDevicePropertyByNameDefault with default headers values
func NewPatchDevicePropertyByNameDefault(code int) *PatchDevicePropertyByNameDefault {
	return &PatchDevicePropertyByNameDefault{
		_statusCode: code,
	}
}

/*PatchDevicePropertyByNameDefault handles this case with default header values.

Error
*/
type PatchDevicePropertyByNameDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch device property by name default response
func (o *PatchDevicePropertyByNameDefault) Code() int {
	return o._statusCode
}

func (o *PatchDevicePropertyByNameDefault) Error() string {
	return fmt.Sprintf("[PATCH /device/devices/{deviceId}/properties/{name}][%d] patchDevicePropertyByName default  %+v", o._statusCode, o.Payload)
}

func (o *PatchDevicePropertyByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
