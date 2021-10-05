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

// GetTeacherByIDReader is a Reader for the GetTeacherByID structure.
type GetTeacherByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTeacherByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTeacherByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTeacherByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTeacherByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTeacherByIDOK creates a GetTeacherByIDOK with default headers values
func NewGetTeacherByIDOK() *GetTeacherByIDOK {
	return &GetTeacherByIDOK{}
}

/* GetTeacherByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTeacherByIDOK struct {
	Payload *models.Teacher
}

func (o *GetTeacherByIDOK) Error() string {
	return fmt.Sprintf("[GET /teacher/{teacherId}][%d] getTeacherByIdOK  %+v", 200, o.Payload)
}
func (o *GetTeacherByIDOK) GetPayload() *models.Teacher {
	return o.Payload
}

func (o *GetTeacherByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Teacher)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeacherByIDBadRequest creates a GetTeacherByIDBadRequest with default headers values
func NewGetTeacherByIDBadRequest() *GetTeacherByIDBadRequest {
	return &GetTeacherByIDBadRequest{}
}

/* GetTeacherByIDBadRequest describes a response with status code 400, with default header values.

Invalid ID supplied
*/
type GetTeacherByIDBadRequest struct {
}

func (o *GetTeacherByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /teacher/{teacherId}][%d] getTeacherByIdBadRequest ", 400)
}

func (o *GetTeacherByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTeacherByIDNotFound creates a GetTeacherByIDNotFound with default headers values
func NewGetTeacherByIDNotFound() *GetTeacherByIDNotFound {
	return &GetTeacherByIDNotFound{}
}

/* GetTeacherByIDNotFound describes a response with status code 404, with default header values.

Teacher not found
*/
type GetTeacherByIDNotFound struct {
}

func (o *GetTeacherByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /teacher/{teacherId}][%d] getTeacherByIdNotFound ", 404)
}

func (o *GetTeacherByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}