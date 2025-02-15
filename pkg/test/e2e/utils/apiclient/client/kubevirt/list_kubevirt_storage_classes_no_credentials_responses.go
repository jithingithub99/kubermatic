// Code generated by go-swagger; DO NOT EDIT.

package kubevirt

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListKubevirtStorageClassesNoCredentialsReader is a Reader for the ListKubevirtStorageClassesNoCredentials structure.
type ListKubevirtStorageClassesNoCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListKubevirtStorageClassesNoCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListKubevirtStorageClassesNoCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListKubevirtStorageClassesNoCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListKubevirtStorageClassesNoCredentialsOK creates a ListKubevirtStorageClassesNoCredentialsOK with default headers values
func NewListKubevirtStorageClassesNoCredentialsOK() *ListKubevirtStorageClassesNoCredentialsOK {
	return &ListKubevirtStorageClassesNoCredentialsOK{}
}

/*
ListKubevirtStorageClassesNoCredentialsOK describes a response with status code 200, with default header values.

StorageClassList
*/
type ListKubevirtStorageClassesNoCredentialsOK struct {
	Payload models.StorageClassList
}

// IsSuccess returns true when this list kubevirt storage classes no credentials o k response has a 2xx status code
func (o *ListKubevirtStorageClassesNoCredentialsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list kubevirt storage classes no credentials o k response has a 3xx status code
func (o *ListKubevirtStorageClassesNoCredentialsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list kubevirt storage classes no credentials o k response has a 4xx status code
func (o *ListKubevirtStorageClassesNoCredentialsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list kubevirt storage classes no credentials o k response has a 5xx status code
func (o *ListKubevirtStorageClassesNoCredentialsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list kubevirt storage classes no credentials o k response a status code equal to that given
func (o *ListKubevirtStorageClassesNoCredentialsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListKubevirtStorageClassesNoCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/kubevirt/storageclasses][%d] listKubevirtStorageClassesNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListKubevirtStorageClassesNoCredentialsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/kubevirt/storageclasses][%d] listKubevirtStorageClassesNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListKubevirtStorageClassesNoCredentialsOK) GetPayload() models.StorageClassList {
	return o.Payload
}

func (o *ListKubevirtStorageClassesNoCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListKubevirtStorageClassesNoCredentialsDefault creates a ListKubevirtStorageClassesNoCredentialsDefault with default headers values
func NewListKubevirtStorageClassesNoCredentialsDefault(code int) *ListKubevirtStorageClassesNoCredentialsDefault {
	return &ListKubevirtStorageClassesNoCredentialsDefault{
		_statusCode: code,
	}
}

/*
ListKubevirtStorageClassesNoCredentialsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListKubevirtStorageClassesNoCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list kubevirt storage classes no credentials default response
func (o *ListKubevirtStorageClassesNoCredentialsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list kubevirt storage classes no credentials default response has a 2xx status code
func (o *ListKubevirtStorageClassesNoCredentialsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list kubevirt storage classes no credentials default response has a 3xx status code
func (o *ListKubevirtStorageClassesNoCredentialsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list kubevirt storage classes no credentials default response has a 4xx status code
func (o *ListKubevirtStorageClassesNoCredentialsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list kubevirt storage classes no credentials default response has a 5xx status code
func (o *ListKubevirtStorageClassesNoCredentialsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list kubevirt storage classes no credentials default response a status code equal to that given
func (o *ListKubevirtStorageClassesNoCredentialsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListKubevirtStorageClassesNoCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/kubevirt/storageclasses][%d] listKubevirtStorageClassesNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListKubevirtStorageClassesNoCredentialsDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/kubevirt/storageclasses][%d] listKubevirtStorageClassesNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListKubevirtStorageClassesNoCredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListKubevirtStorageClassesNoCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
