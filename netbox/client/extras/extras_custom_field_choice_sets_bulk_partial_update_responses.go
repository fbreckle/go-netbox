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

// ExtrasCustomFieldChoiceSetsBulkPartialUpdateReader is a Reader for the ExtrasCustomFieldChoiceSetsBulkPartialUpdate structure.
type ExtrasCustomFieldChoiceSetsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomFieldChoiceSetsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasCustomFieldChoiceSetsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasCustomFieldChoiceSetsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasCustomFieldChoiceSetsBulkPartialUpdateOK creates a ExtrasCustomFieldChoiceSetsBulkPartialUpdateOK with default headers values
func NewExtrasCustomFieldChoiceSetsBulkPartialUpdateOK() *ExtrasCustomFieldChoiceSetsBulkPartialUpdateOK {
	return &ExtrasCustomFieldChoiceSetsBulkPartialUpdateOK{}
}

/*
ExtrasCustomFieldChoiceSetsBulkPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasCustomFieldChoiceSetsBulkPartialUpdateOK extras custom field choice sets bulk partial update o k
*/
type ExtrasCustomFieldChoiceSetsBulkPartialUpdateOK struct {
	Payload *models.CustomFieldChoiceSet
}

// IsSuccess returns true when this extras custom field choice sets bulk partial update o k response has a 2xx status code
func (o *ExtrasCustomFieldChoiceSetsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras custom field choice sets bulk partial update o k response has a 3xx status code
func (o *ExtrasCustomFieldChoiceSetsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras custom field choice sets bulk partial update o k response has a 4xx status code
func (o *ExtrasCustomFieldChoiceSetsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras custom field choice sets bulk partial update o k response has a 5xx status code
func (o *ExtrasCustomFieldChoiceSetsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras custom field choice sets bulk partial update o k response a status code equal to that given
func (o *ExtrasCustomFieldChoiceSetsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExtrasCustomFieldChoiceSetsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/custom-field-choice-sets/][%d] extrasCustomFieldChoiceSetsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasCustomFieldChoiceSetsBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /extras/custom-field-choice-sets/][%d] extrasCustomFieldChoiceSetsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasCustomFieldChoiceSetsBulkPartialUpdateOK) GetPayload() *models.CustomFieldChoiceSet {
	return o.Payload
}

func (o *ExtrasCustomFieldChoiceSetsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomFieldChoiceSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasCustomFieldChoiceSetsBulkPartialUpdateDefault creates a ExtrasCustomFieldChoiceSetsBulkPartialUpdateDefault with default headers values
func NewExtrasCustomFieldChoiceSetsBulkPartialUpdateDefault(code int) *ExtrasCustomFieldChoiceSetsBulkPartialUpdateDefault {
	return &ExtrasCustomFieldChoiceSetsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
ExtrasCustomFieldChoiceSetsBulkPartialUpdateDefault describes a response with status code -1, with default header values.

ExtrasCustomFieldChoiceSetsBulkPartialUpdateDefault extras custom field choice sets bulk partial update default
*/
type ExtrasCustomFieldChoiceSetsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras custom field choice sets bulk partial update default response
func (o *ExtrasCustomFieldChoiceSetsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras custom field choice sets bulk partial update default response has a 2xx status code
func (o *ExtrasCustomFieldChoiceSetsBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras custom field choice sets bulk partial update default response has a 3xx status code
func (o *ExtrasCustomFieldChoiceSetsBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras custom field choice sets bulk partial update default response has a 4xx status code
func (o *ExtrasCustomFieldChoiceSetsBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras custom field choice sets bulk partial update default response has a 5xx status code
func (o *ExtrasCustomFieldChoiceSetsBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras custom field choice sets bulk partial update default response a status code equal to that given
func (o *ExtrasCustomFieldChoiceSetsBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasCustomFieldChoiceSetsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /extras/custom-field-choice-sets/][%d] extras_custom-field-choice-sets_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasCustomFieldChoiceSetsBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /extras/custom-field-choice-sets/][%d] extras_custom-field-choice-sets_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasCustomFieldChoiceSetsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasCustomFieldChoiceSetsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
