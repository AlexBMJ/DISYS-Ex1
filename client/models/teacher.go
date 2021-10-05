// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Teacher teacher
// Example: {"courses":[{"id":6,"name":"name"},{"id":6,"name":"name"}],"email":"jodo@itu.dk","firstName":"John","id":0,"lastName":"Doe","password":"password","phone":"98765432","photoUrls":["photoUrls","photoUrls"],"status":"teaching","username":"jodo"}
//
// swagger:model Teacher
type Teacher struct {

	// courses
	Courses []*Course `json:"courses"`

	// email
	// Example: jodo@itu.dk
	Email string `json:"email,omitempty"`

	// first name
	// Example: John
	FirstName string `json:"firstName,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// last name
	// Example: Doe
	LastName string `json:"lastName,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// phone
	// Example: 98765432
	Phone string `json:"phone,omitempty"`

	// photo urls
	PhotoUrls []string `json:"photoUrls" xml:"photoUrl"`

	// teacher status
	// Enum: [teaching on_leave retired]
	Status string `json:"status,omitempty"`

	// username
	// Example: jodo
	Username string `json:"username,omitempty"`
}

// Validate validates this teacher
func (m *Teacher) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCourses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Teacher) validateCourses(formats strfmt.Registry) error {
	if swag.IsZero(m.Courses) { // not required
		return nil
	}

	for i := 0; i < len(m.Courses); i++ {
		if swag.IsZero(m.Courses[i]) { // not required
			continue
		}

		if m.Courses[i] != nil {
			if err := m.Courses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("courses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var teacherTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["teaching","on_leave","retired"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		teacherTypeStatusPropEnum = append(teacherTypeStatusPropEnum, v)
	}
}

const (

	// TeacherStatusTeaching captures enum value "teaching"
	TeacherStatusTeaching string = "teaching"

	// TeacherStatusOnLeave captures enum value "on_leave"
	TeacherStatusOnLeave string = "on_leave"

	// TeacherStatusRetired captures enum value "retired"
	TeacherStatusRetired string = "retired"
)

// prop value enum
func (m *Teacher) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, teacherTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Teacher) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this teacher based on the context it is used
func (m *Teacher) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCourses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Teacher) contextValidateCourses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Courses); i++ {

		if m.Courses[i] != nil {
			if err := m.Courses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("courses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Teacher) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Teacher) UnmarshalBinary(b []byte) error {
	var res Teacher
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
