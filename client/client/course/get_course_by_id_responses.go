// Code generated by go-swagger; DO NOT EDIT.

package course

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"example.com/restclient/models"
)

// GetCourseByIDReader is a Reader for the GetCourseByID structure.
type GetCourseByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCourseByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCourseByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCourseByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCourseByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCourseByIDOK creates a GetCourseByIDOK with default headers values
func NewGetCourseByIDOK() *GetCourseByIDOK {
	return &GetCourseByIDOK{}
}

/* GetCourseByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetCourseByIDOK struct {
	Payload *models.Course
}

func (o *GetCourseByIDOK) Error() string {
	return fmt.Sprintf("[GET /course/{courseId}][%d] getCourseByIdOK  %+v", 200, o.Payload)
}
func (o *GetCourseByIDOK) GetPayload() *models.Course {
	return o.Payload
}

func (o *GetCourseByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Course)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCourseByIDBadRequest creates a GetCourseByIDBadRequest with default headers values
func NewGetCourseByIDBadRequest() *GetCourseByIDBadRequest {
	return &GetCourseByIDBadRequest{}
}

/* GetCourseByIDBadRequest describes a response with status code 400, with default header values.

Invalid ID supplied
*/
type GetCourseByIDBadRequest struct {
}

func (o *GetCourseByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /course/{courseId}][%d] getCourseByIdBadRequest ", 400)
}

func (o *GetCourseByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCourseByIDNotFound creates a GetCourseByIDNotFound with default headers values
func NewGetCourseByIDNotFound() *GetCourseByIDNotFound {
	return &GetCourseByIDNotFound{}
}

/* GetCourseByIDNotFound describes a response with status code 404, with default header values.

Course not found
*/
type GetCourseByIDNotFound struct {
}

func (o *GetCourseByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /course/{courseId}][%d] getCourseByIdNotFound ", 404)
}

func (o *GetCourseByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
