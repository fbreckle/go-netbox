// Code generated by go-swagger; DO NOT EDIT.

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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// ExtrasConfigTemplatesPartialUpdateReader is a Reader for the ExtrasConfigTemplatesPartialUpdate structure.
type ExtrasConfigTemplatesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasConfigTemplatesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasConfigTemplatesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasConfigTemplatesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasConfigTemplatesPartialUpdateOK creates a ExtrasConfigTemplatesPartialUpdateOK with default headers values
func NewExtrasConfigTemplatesPartialUpdateOK() *ExtrasConfigTemplatesPartialUpdateOK {
	return &ExtrasConfigTemplatesPartialUpdateOK{}
}

/*
ExtrasConfigTemplatesPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasConfigTemplatesPartialUpdateOK extras config templates partial update o k
*/
type ExtrasConfigTemplatesPartialUpdateOK struct {
	Payload *models.ConfigTemplate
}

// IsSuccess returns true when this extras config templates partial update o k response has a 2xx status code
func (o *ExtrasConfigTemplatesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras config templates partial update o k response has a 3xx status code
func (o *ExtrasConfigTemplatesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras config templates partial update o k response has a 4xx status code
func (o *ExtrasConfigTemplatesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras config templates partial update o k response has a 5xx status code
func (o *ExtrasConfigTemplatesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras config templates partial update o k response a status code equal to that given
func (o *ExtrasConfigTemplatesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras config templates partial update o k response
func (o *ExtrasConfigTemplatesPartialUpdateOK) Code() int {
	return 200
}

func (o *ExtrasConfigTemplatesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/config-templates/{id}/][%d] extrasConfigTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasConfigTemplatesPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /extras/config-templates/{id}/][%d] extrasConfigTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasConfigTemplatesPartialUpdateOK) GetPayload() *models.ConfigTemplate {
	return o.Payload
}

func (o *ExtrasConfigTemplatesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasConfigTemplatesPartialUpdateDefault creates a ExtrasConfigTemplatesPartialUpdateDefault with default headers values
func NewExtrasConfigTemplatesPartialUpdateDefault(code int) *ExtrasConfigTemplatesPartialUpdateDefault {
	return &ExtrasConfigTemplatesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
ExtrasConfigTemplatesPartialUpdateDefault describes a response with status code -1, with default header values.

ExtrasConfigTemplatesPartialUpdateDefault extras config templates partial update default
*/
type ExtrasConfigTemplatesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this extras config templates partial update default response has a 2xx status code
func (o *ExtrasConfigTemplatesPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras config templates partial update default response has a 3xx status code
func (o *ExtrasConfigTemplatesPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras config templates partial update default response has a 4xx status code
func (o *ExtrasConfigTemplatesPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras config templates partial update default response has a 5xx status code
func (o *ExtrasConfigTemplatesPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras config templates partial update default response a status code equal to that given
func (o *ExtrasConfigTemplatesPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the extras config templates partial update default response
func (o *ExtrasConfigTemplatesPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasConfigTemplatesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /extras/config-templates/{id}/][%d] extras_config-templates_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasConfigTemplatesPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /extras/config-templates/{id}/][%d] extras_config-templates_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasConfigTemplatesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasConfigTemplatesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}