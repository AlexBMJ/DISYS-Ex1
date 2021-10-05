// Code generated by go-swagger; DO NOT EDIT.

package student

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new student API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for student API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddStudent(params *AddStudentParams, opts ...ClientOption) (*AddStudentOK, error)

	DeleteStudent(params *DeleteStudentParams, opts ...ClientOption) error

	FindStudentByStatus(params *FindStudentByStatusParams, opts ...ClientOption) (*FindStudentByStatusOK, error)

	GetStudentByID(params *GetStudentByIDParams, opts ...ClientOption) (*GetStudentByIDOK, error)

	UpdateStudent(params *UpdateStudentParams, opts ...ClientOption) (*UpdateStudentOK, error)

	UploadStudentPhoto(params *UploadStudentPhotoParams, opts ...ClientOption) (*UploadStudentPhotoOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddStudent adds a new student

  Student at ITU
*/
func (a *Client) AddStudent(params *AddStudentParams, opts ...ClientOption) (*AddStudentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddStudentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addStudent",
		Method:             "POST",
		PathPattern:        "/student",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddStudentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddStudentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addStudent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteStudent deletes a student
*/
func (a *Client) DeleteStudent(params *DeleteStudentParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteStudentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteStudent",
		Method:             "DELETE",
		PathPattern:        "/student/{studentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteStudentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
  FindStudentByStatus finds student by status

  Multiple status values can be provided with comma separated strings
*/
func (a *Client) FindStudentByStatus(params *FindStudentByStatusParams, opts ...ClientOption) (*FindStudentByStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindStudentByStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findStudentByStatus",
		Method:             "GET",
		PathPattern:        "/student/findByStatus",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FindStudentByStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindStudentByStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findStudentByStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetStudentByID finds student by ID

  Returns a single student
*/
func (a *Client) GetStudentByID(params *GetStudentByIDParams, opts ...ClientOption) (*GetStudentByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStudentByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getStudentById",
		Method:             "GET",
		PathPattern:        "/student/{studentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetStudentByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStudentByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getStudentById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateStudent updates a student
*/
func (a *Client) UpdateStudent(params *UpdateStudentParams, opts ...ClientOption) (*UpdateStudentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateStudentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateStudent",
		Method:             "PUT",
		PathPattern:        "/student",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateStudentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateStudentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateStudent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UploadStudentPhoto uploads an image
*/
func (a *Client) UploadStudentPhoto(params *UploadStudentPhotoParams, opts ...ClientOption) (*UploadStudentPhotoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadStudentPhotoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "uploadStudentPhoto",
		Method:             "POST",
		PathPattern:        "/student/{studentId}/uploadImage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UploadStudentPhotoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UploadStudentPhotoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for uploadStudentPhoto: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}