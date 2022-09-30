// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/odpf/dex/generated/models"
)

// ListFirehosesReader is a Reader for the ListFirehoses structure.
type ListFirehosesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListFirehosesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListFirehosesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewListFirehosesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListFirehosesOK creates a ListFirehosesOK with default headers values
func NewListFirehosesOK() *ListFirehosesOK {
	return &ListFirehosesOK{}
}

/*
ListFirehosesOK describes a response with status code 200, with default header values.

successful operation
*/
type ListFirehosesOK struct {
	Payload *models.FirehoseArray
}

// IsSuccess returns true when this list firehoses o k response has a 2xx status code
func (o *ListFirehosesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list firehoses o k response has a 3xx status code
func (o *ListFirehosesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list firehoses o k response has a 4xx status code
func (o *ListFirehosesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list firehoses o k response has a 5xx status code
func (o *ListFirehosesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list firehoses o k response a status code equal to that given
func (o *ListFirehosesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListFirehosesOK) Error() string {
	return fmt.Sprintf("[GET /projects/{projectId}/firehoses][%d] listFirehosesOK  %+v", 200, o.Payload)
}

func (o *ListFirehosesOK) String() string {
	return fmt.Sprintf("[GET /projects/{projectId}/firehoses][%d] listFirehosesOK  %+v", 200, o.Payload)
}

func (o *ListFirehosesOK) GetPayload() *models.FirehoseArray {
	return o.Payload
}

func (o *ListFirehosesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FirehoseArray)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFirehosesInternalServerError creates a ListFirehosesInternalServerError with default headers values
func NewListFirehosesInternalServerError() *ListFirehosesInternalServerError {
	return &ListFirehosesInternalServerError{}
}

/*
ListFirehosesInternalServerError describes a response with status code 500, with default header values.

internal error
*/
type ListFirehosesInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list firehoses internal server error response has a 2xx status code
func (o *ListFirehosesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list firehoses internal server error response has a 3xx status code
func (o *ListFirehosesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list firehoses internal server error response has a 4xx status code
func (o *ListFirehosesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list firehoses internal server error response has a 5xx status code
func (o *ListFirehosesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list firehoses internal server error response a status code equal to that given
func (o *ListFirehosesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ListFirehosesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{projectId}/firehoses][%d] listFirehosesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListFirehosesInternalServerError) String() string {
	return fmt.Sprintf("[GET /projects/{projectId}/firehoses][%d] listFirehosesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListFirehosesInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListFirehosesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}