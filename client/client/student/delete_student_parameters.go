// Code generated by go-swagger; DO NOT EDIT.

package student

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeleteStudentParams creates a new DeleteStudentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteStudentParams() *DeleteStudentParams {
	return &DeleteStudentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteStudentParamsWithTimeout creates a new DeleteStudentParams object
// with the ability to set a timeout on a request.
func NewDeleteStudentParamsWithTimeout(timeout time.Duration) *DeleteStudentParams {
	return &DeleteStudentParams{
		timeout: timeout,
	}
}

// NewDeleteStudentParamsWithContext creates a new DeleteStudentParams object
// with the ability to set a context for a request.
func NewDeleteStudentParamsWithContext(ctx context.Context) *DeleteStudentParams {
	return &DeleteStudentParams{
		Context: ctx,
	}
}

// NewDeleteStudentParamsWithHTTPClient creates a new DeleteStudentParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteStudentParamsWithHTTPClient(client *http.Client) *DeleteStudentParams {
	return &DeleteStudentParams{
		HTTPClient: client,
	}
}

/* DeleteStudentParams contains all the parameters to send to the API endpoint
   for the delete student operation.

   Typically these are written to a http.Request.
*/
type DeleteStudentParams struct {

	/* StudentID.

	   Id of student to delete

	   Format: int64
	*/
	StudentID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete student params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteStudentParams) WithDefaults() *DeleteStudentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete student params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteStudentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete student params
func (o *DeleteStudentParams) WithTimeout(timeout time.Duration) *DeleteStudentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete student params
func (o *DeleteStudentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete student params
func (o *DeleteStudentParams) WithContext(ctx context.Context) *DeleteStudentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete student params
func (o *DeleteStudentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete student params
func (o *DeleteStudentParams) WithHTTPClient(client *http.Client) *DeleteStudentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete student params
func (o *DeleteStudentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStudentID adds the studentID to the delete student params
func (o *DeleteStudentParams) WithStudentID(studentID int64) *DeleteStudentParams {
	o.SetStudentID(studentID)
	return o
}

// SetStudentID adds the studentId to the delete student params
func (o *DeleteStudentParams) SetStudentID(studentID int64) {
	o.StudentID = studentID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteStudentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param studentId
	if err := r.SetPathParam("studentId", swag.FormatInt64(o.StudentID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
