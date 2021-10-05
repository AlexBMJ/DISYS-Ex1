/*
 * ITU API
 *
 * ITU REST API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Teacher struct {

	Id int64 `json:"id,omitempty"`

	Username string `json:"username,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	LastName string `json:"lastName,omitempty"`

	Email string `json:"email,omitempty"`

	Password string `json:"password,omitempty"`

	Phone string `json:"phone,omitempty"`

	Courses []Course `json:"courses,omitempty"`

	PhotoUrls []string `json:"photoUrls,omitempty"`

	// teacher status
	Status string `json:"status,omitempty"`
}