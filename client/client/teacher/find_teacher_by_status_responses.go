// Code generated by go-swagger; DO NOT EDIT.

package teacher

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"example.com/restclient/models"
)

// FindTeacherByStatusReader is a Reader for the FindTeacherByStatus structure.
type FindTeacherByStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindTeacherByStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindTeacherByStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFindTeacherByStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFindTeacherByStatusOK creates a FindTeacherByStatusOK with default headers values
func NewFindTeacherByStatusOK() *FindTeacherByStatusOK {
	return &FindTeacherByStatusOK{}
}

/* FindTeacherByStatusOK describes a response with status code 200, with default header values.

Success
*/
type FindTeacherByStatusOK struct {
	Payload []*models.Teacher
}

func (o *FindTeacherByStatusOK) Error() string {
	return fmt.Sprintf("[GET /teacher/findByStatus][%d] findTeacherByStatusOK  %+v", 200, o.Payload)
}
func (o *FindTeacherByStatusOK) GetPayload() []*models.Teacher {
	return o.Payload
}

func (o *FindTeacherByStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindTeacherByStatusBadRequest creates a FindTeacherByStatusBadRequest with default headers values
func NewFindTeacherByStatusBadRequest() *FindTeacherByStatusBadRequest {
	return &FindTeacherByStatusBadRequest{}
}

/* FindTeacherByStatusBadRequest describes a response with status code 400, with default header values.

Invalid status value
*/
type FindTeacherByStatusBadRequest struct {
}

func (o *FindTeacherByStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /teacher/findByStatus][%d] findTeacherByStatusBadRequest ", 400)
}

func (o *FindTeacherByStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}