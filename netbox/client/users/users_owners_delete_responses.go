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
)

// UsersOwnersDeleteReader is a Reader for the UsersOwnersDelete structure.
type UsersOwnersDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersOwnersDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUsersOwnersDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersOwnersDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersOwnersDeleteNoContent creates a UsersOwnersDeleteNoContent with default headers values
func NewUsersOwnersDeleteNoContent() *UsersOwnersDeleteNoContent {
	return &UsersOwnersDeleteNoContent{}
}

/*
UsersOwnersDeleteNoContent describes a response with status code 204, with default header values.

UsersOwnersDeleteNoContent users owners delete no content
*/
type UsersOwnersDeleteNoContent struct {
}

// IsSuccess returns true when this users owners delete no content response has a 2xx status code
func (o *UsersOwnersDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users owners delete no content response has a 3xx status code
func (o *UsersOwnersDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users owners delete no content response has a 4xx status code
func (o *UsersOwnersDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this users owners delete no content response has a 5xx status code
func (o *UsersOwnersDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this users owners delete no content response a status code equal to that given
func (o *UsersOwnersDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the users owners delete no content response
func (o *UsersOwnersDeleteNoContent) Code() int {
	return 204
}

func (o *UsersOwnersDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/owners/{id}/][%d] usersOwnersDeleteNoContent", 204)
}

func (o *UsersOwnersDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /users/owners/{id}/][%d] usersOwnersDeleteNoContent", 204)
}

func (o *UsersOwnersDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUsersOwnersDeleteDefault creates a UsersOwnersDeleteDefault with default headers values
func NewUsersOwnersDeleteDefault(code int) *UsersOwnersDeleteDefault {
	return &UsersOwnersDeleteDefault{
		_statusCode: code,
	}
}

/*
UsersOwnersDeleteDefault describes a response with status code -1, with default header values.

UsersOwnersDeleteDefault users owners delete default
*/
type UsersOwnersDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this users owners delete default response has a 2xx status code
func (o *UsersOwnersDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this users owners delete default response has a 3xx status code
func (o *UsersOwnersDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this users owners delete default response has a 4xx status code
func (o *UsersOwnersDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this users owners delete default response has a 5xx status code
func (o *UsersOwnersDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this users owners delete default response a status code equal to that given
func (o *UsersOwnersDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the users owners delete default response
func (o *UsersOwnersDeleteDefault) Code() int {
	return o._statusCode
}

func (o *UsersOwnersDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /users/owners/{id}/][%d] users_owners_delete default %s", o._statusCode, payload)
}

func (o *UsersOwnersDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /users/owners/{id}/][%d] users_owners_delete default %s", o._statusCode, payload)
}

func (o *UsersOwnersDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersOwnersDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
