package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cloudfoundry/bosh-softlayer-cpi/softlayer/pool/models"
)

// FindVmsByFiltersReader is a Reader for the FindVmsByFilters structure.
type FindVmsByFiltersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindVmsByFiltersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFindVmsByFiltersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewFindVmsByFiltersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewFindVmsByFiltersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewFindVmsByFiltersOK creates a FindVmsByFiltersOK with default headers values
func NewFindVmsByFiltersOK() *FindVmsByFiltersOK {
	return &FindVmsByFiltersOK{}
}

/*FindVmsByFiltersOK handles this case with default header values.

successful operation
*/
type FindVmsByFiltersOK struct {
	Payload *models.VmsResponse
}

func (o *FindVmsByFiltersOK) Error() string {
	return fmt.Sprintf("[POST /vms/findByFilters][%d] findVmsByFiltersOK  %+v", 200, o.Payload)
}

func (o *FindVmsByFiltersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VmsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindVmsByFiltersNotFound creates a FindVmsByFiltersNotFound with default headers values
func NewFindVmsByFiltersNotFound() *FindVmsByFiltersNotFound {
	return &FindVmsByFiltersNotFound{}
}

/*FindVmsByFiltersNotFound handles this case with default header values.

vm not found
*/
type FindVmsByFiltersNotFound struct {
}

func (o *FindVmsByFiltersNotFound) Error() string {
	return fmt.Sprintf("[POST /vms/findByFilters][%d] findVmsByFiltersNotFound ", 404)
}

func (o *FindVmsByFiltersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindVmsByFiltersDefault creates a FindVmsByFiltersDefault with default headers values
func NewFindVmsByFiltersDefault(code int) *FindVmsByFiltersDefault {
	return &FindVmsByFiltersDefault{
		_statusCode: code,
	}
}

/*FindVmsByFiltersDefault handles this case with default header values.

unexpected error
*/
type FindVmsByFiltersDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the find vms by filters default response
func (o *FindVmsByFiltersDefault) Code() int {
	return o._statusCode
}

func (o *FindVmsByFiltersDefault) Error() string {
	return fmt.Sprintf("[POST /vms/findByFilters][%d] findVmsByFilters default  %+v", o._statusCode, o.Payload)
}

func (o *FindVmsByFiltersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}