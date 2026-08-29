// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package users

// This file was manually added to extend the client with NetBox's Owner/OwnerGroup
// endpoints (netbox-community/netbox users app), which are not yet covered by the
// upstream swagger-generated client.

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// UsersOwnerGroupsPartialUpdateReader is a Reader for the UsersOwnerGroupsPartialUpdate structure.
type UsersOwnerGroupsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersOwnerGroupsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersOwnerGroupsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersOwnerGroupsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersOwnerGroupsPartialUpdateOK creates a UsersOwnerGroupsPartialUpdateOK with default headers values
func NewUsersOwnerGroupsPartialUpdateOK() *UsersOwnerGroupsPartialUpdateOK {
	return &UsersOwnerGroupsPartialUpdateOK{}
}

/*
UsersOwnerGroupsPartialUpdateOK describes a response with status code 200, with default header values.

UsersOwnerGroupsPartialUpdateOK users owner groups partial update o k
*/
type UsersOwnerGroupsPartialUpdateOK struct {
	Payload *models.OwnerGroup
}

// IsSuccess returns true when this users owner groups partial update o k response has a 2xx status code
func (o *UsersOwnerGroupsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users owner groups partial update o k response has a 3xx status code
func (o *UsersOwnerGroupsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users owner groups partial update o k response has a 4xx status code
func (o *UsersOwnerGroupsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users owner groups partial update o k response has a 5xx status code
func (o *UsersOwnerGroupsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users owner groups partial update o k response a status code equal to that given
func (o *UsersOwnerGroupsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the users owner groups partial update o k response
func (o *UsersOwnerGroupsPartialUpdateOK) Code() int {
	return 200
}

func (o *UsersOwnerGroupsPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /users/owner-groups/{id}/][%d] usersOwnerGroupsPartialUpdateOK %s", 200, payload)
}

func (o *UsersOwnerGroupsPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /users/owner-groups/{id}/][%d] usersOwnerGroupsPartialUpdateOK %s", 200, payload)
}

func (o *UsersOwnerGroupsPartialUpdateOK) GetPayload() *models.OwnerGroup {
	return o.Payload
}

func (o *UsersOwnerGroupsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OwnerGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersOwnerGroupsPartialUpdateDefault creates a UsersOwnerGroupsPartialUpdateDefault with default headers values
func NewUsersOwnerGroupsPartialUpdateDefault(code int) *UsersOwnerGroupsPartialUpdateDefault {
	return &UsersOwnerGroupsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
UsersOwnerGroupsPartialUpdateDefault describes a response with status code -1, with default header values.

UsersOwnerGroupsPartialUpdateDefault users owner groups partial update default
*/
type UsersOwnerGroupsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this users owner groups partial update default response has a 2xx status code
func (o *UsersOwnerGroupsPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this users owner groups partial update default response has a 3xx status code
func (o *UsersOwnerGroupsPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this users owner groups partial update default response has a 4xx status code
func (o *UsersOwnerGroupsPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this users owner groups partial update default response has a 5xx status code
func (o *UsersOwnerGroupsPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this users owner groups partial update default response a status code equal to that given
func (o *UsersOwnerGroupsPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the users owner groups partial update default response
func (o *UsersOwnerGroupsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *UsersOwnerGroupsPartialUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /users/owner-groups/{id}/][%d] users_owner_groups_partial_update default %s", o._statusCode, payload)
}

func (o *UsersOwnerGroupsPartialUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /users/owner-groups/{id}/][%d] users_owner_groups_partial_update default %s", o._statusCode, payload)
}

func (o *UsersOwnerGroupsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersOwnerGroupsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
