// Code generated by go-swagger; DO NOT EDIT.

package datacenter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// DeleteDCReader is a Reader for the DeleteDC structure.
type DeleteDCReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDCReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDCOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteDCUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteDCForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteDCDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteDCOK creates a DeleteDCOK with default headers values
func NewDeleteDCOK() *DeleteDCOK {
	return &DeleteDCOK{}
}

/*
DeleteDCOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type DeleteDCOK struct {
}

// IsSuccess returns true when this delete d c o k response has a 2xx status code
func (o *DeleteDCOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete d c o k response has a 3xx status code
func (o *DeleteDCOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete d c o k response has a 4xx status code
func (o *DeleteDCOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete d c o k response has a 5xx status code
func (o *DeleteDCOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete d c o k response a status code equal to that given
func (o *DeleteDCOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteDCOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/seed/{seed_name}/dc/{dc}][%d] deleteDCOK ", 200)
}

func (o *DeleteDCOK) String() string {
	return fmt.Sprintf("[DELETE /api/v1/seed/{seed_name}/dc/{dc}][%d] deleteDCOK ", 200)
}

func (o *DeleteDCOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDCUnauthorized creates a DeleteDCUnauthorized with default headers values
func NewDeleteDCUnauthorized() *DeleteDCUnauthorized {
	return &DeleteDCUnauthorized{}
}

/*
DeleteDCUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type DeleteDCUnauthorized struct {
}

// IsSuccess returns true when this delete d c unauthorized response has a 2xx status code
func (o *DeleteDCUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete d c unauthorized response has a 3xx status code
func (o *DeleteDCUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete d c unauthorized response has a 4xx status code
func (o *DeleteDCUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete d c unauthorized response has a 5xx status code
func (o *DeleteDCUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete d c unauthorized response a status code equal to that given
func (o *DeleteDCUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteDCUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/seed/{seed_name}/dc/{dc}][%d] deleteDCUnauthorized ", 401)
}

func (o *DeleteDCUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v1/seed/{seed_name}/dc/{dc}][%d] deleteDCUnauthorized ", 401)
}

func (o *DeleteDCUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDCForbidden creates a DeleteDCForbidden with default headers values
func NewDeleteDCForbidden() *DeleteDCForbidden {
	return &DeleteDCForbidden{}
}

/*
DeleteDCForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type DeleteDCForbidden struct {
}

// IsSuccess returns true when this delete d c forbidden response has a 2xx status code
func (o *DeleteDCForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete d c forbidden response has a 3xx status code
func (o *DeleteDCForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete d c forbidden response has a 4xx status code
func (o *DeleteDCForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete d c forbidden response has a 5xx status code
func (o *DeleteDCForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete d c forbidden response a status code equal to that given
func (o *DeleteDCForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteDCForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/seed/{seed_name}/dc/{dc}][%d] deleteDCForbidden ", 403)
}

func (o *DeleteDCForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v1/seed/{seed_name}/dc/{dc}][%d] deleteDCForbidden ", 403)
}

func (o *DeleteDCForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDCDefault creates a DeleteDCDefault with default headers values
func NewDeleteDCDefault(code int) *DeleteDCDefault {
	return &DeleteDCDefault{
		_statusCode: code,
	}
}

/*
DeleteDCDefault describes a response with status code -1, with default header values.

errorResponse
*/
type DeleteDCDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete d c default response
func (o *DeleteDCDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete d c default response has a 2xx status code
func (o *DeleteDCDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete d c default response has a 3xx status code
func (o *DeleteDCDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete d c default response has a 4xx status code
func (o *DeleteDCDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete d c default response has a 5xx status code
func (o *DeleteDCDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete d c default response a status code equal to that given
func (o *DeleteDCDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteDCDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/seed/{seed_name}/dc/{dc}][%d] deleteDC default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteDCDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v1/seed/{seed_name}/dc/{dc}][%d] deleteDC default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteDCDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteDCDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
